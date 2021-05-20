package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type List struct {
	ID        int      `json:"id"`
	Title     string   `json:"title" validate:"required"`
	Body      string   `json:"body" validate:"required"`
	Tags      []string `json:"tags" validate:"required"`
	CreatedAt int64    `json:"created_at"`
}

type Tag struct {
	ListId int
	Name   string
}

type ResponseError struct {
	Success bool   `json:"sucess"`
	Err     string `json:"err"`
	Message string `json:"message"`
}

func ConnectDB() *sql.DB {
	config := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?parseTime=true", GoDotEnv("MYSQL_USER"), GoDotEnv("MYSQL_PASSWORD"), GoDotEnv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", config)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func getList() ([]List, error) {
	db := ConnectDB()

	rows, err := db.Query("SELECT id, title, body, tags, created_at FROM ads ORDER BY created_at DESC LIMIT 5")
	if err != nil {
		return nil, fmt.Errorf("unable to execute query due: %v", err)
	}
	var lists []List
	for rows.Next() {
		var list List
		var tagsJSON []byte
		err = rows.Scan(&list.ID, &list.Title, &list.Body, &tagsJSON, &list.CreatedAt)
		if err != nil {
			continue
		}
		if len(tagsJSON) > 0 {
			err = json.Unmarshal(tagsJSON, &list.Tags)
			if err != nil {
				continue
			}
		}
		lists = append(lists, list)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to iterate rows due: %v", err)
	}

	return lists, nil
}

func addList(list List) (*List, error) {
	db := ConnectDB()

	tagsJSON, _ := json.Marshal(list.Tags)
	nowTs := time.Now().Unix()

	query := `INSERT INTO ads (title, body, tags, created_at) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(query, list.Title, list.Body, string(tagsJSON), nowTs)
	if err != nil {
		return nil, fmt.Errorf("unable to execute insert due: %v", err)
	}
	recordID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("unable to get last inserted id due: %v", err)
	}

	list.ID = int(recordID)
	list.CreatedAt = nowTs

	return &list, nil
}
