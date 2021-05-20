package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

type StorageConfigs struct {
	DBUser string
	DBPass string
	DBName string
}

func NewStorage(configs StorageConfigs) (*Storage, error) {
	// initialize db object
	config := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?parseTime=true", configs.DBUser, configs.DBPass, configs.DBName)
	d, err := sql.Open("mysql", config)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize db due: %v", err)
	}
	// establish db connection by ping
	err = d.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to ping db due: %v", err)
	}

	return &Storage{db: d}, nil
}

func (s *Storage) GetList() ([]List, error) {
	// execute sql query
	rows, err := s.db.Query("SELECT id, title, body, tags, created_at FROM ads ORDER BY created_at DESC LIMIT 5")
	if err != nil {
		return nil, fmt.Errorf("unable to execute query due: %v", err)
	}
	defer rows.Close()
	// iterate rows
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
	// check error in rows iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to iterate rows due: %v", err)
	}

	return lists, nil
}

func (s *Storage) AddList(list List) (*List, error) {
	// encode tags
	tagsJSON, _ := json.Marshal(list.Tags)
	// get current timestamp for `created_at` value
	nowTs := time.Now().Unix()

	// execute sql query
	query := `INSERT INTO ads (title, body, tags, created_at) VALUES (?, ?, ?, ?)`
	result, err := s.db.Exec(query, list.Title, list.Body, string(tagsJSON), nowTs)
	if err != nil {
		return nil, fmt.Errorf("unable to execute insert due: %v", err)
	}
	// get record id
	recordID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("unable to get last inserted id due: %v", err)
	}
	// assign id & created at value to list
	list.ID = int(recordID)
	list.CreatedAt = nowTs

	return &list, nil
}
