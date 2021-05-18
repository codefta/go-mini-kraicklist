package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type List struct {
	ID int `json:"id"`
	Title string `json:"title" validate:"required"`
	Body string	`json:"body" validate:"required"`
	Tags []string `json:"tags" validate:"required"`
	CreatedAt int64 `json:"created_at"`
}

type Tag struct {
	ListId int
	Name string
}

type ResponseError struct {
	Success bool `json:"sucess"`
	Err string `json:"err"`
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


func getList() []List{
	db := ConnectDB()

	rowLists, err := db.Query("SELECT * FROM list ORDER BY created_at DESC LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}

	rowTags, err := db.Query("SELECT list_id, name tag_name FROM tag_list JOIN tag on tag.id = tag_id")

	var lists []List

	var tags []string

	for rowLists.Next() {
		var mysqlDate string
		var list List
		rowLists.Scan(&list.ID, &list.Title, &list.Body, &mysqlDate)

		for rowTags.Next() {
			var tag Tag
			rowTags.Scan(&tag.ListId, &tag.Name)

			if list.ID == tag.ListId {
				tags = append(tags, tag.Name)
			}
		}

		list.Tags = tags
		t, _ := time.Parse(time.RFC3339, mysqlDate)
		list.CreatedAt = t.Unix()
		
		lists = append(lists, list)
	}

	defer rowLists.Close()
	defer rowTags.Close()
	defer db.Close()

	return lists
}

func addList(list *List) List{
	db := ConnectDB()
	
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	sqlStatement := `INSERT INTO list (title, body) VALUES(?, ?)`
	result, err := db.Exec(sqlStatement, list.Title, list.Body)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	lastInsertIdList, err := result.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	for _, val := range list.Tags {
		var tag Tag
		db.QueryRow(`SELECT id, name from tag where name=? LIMIT 1`, val).Scan(&tag.ListId, &tag.Name)

		var lastInsertIdTag int64
		
		if tag.Name != val {
			sqlStatement = `INSERT INTO tag (name) VALUES(?)`
			result, err = db.Exec(sqlStatement, val)
			if err != nil {
				_ = tx.Rollback()
				log.Fatal(err)
			}

			lastInsertIdTag, err = result.LastInsertId()

			if err != nil {
				_ = tx.Rollback()
				log.Fatal(err)
			}
		} else {
			lastInsertIdTag = int64(tag.ListId)
		}

		_, err = db.Exec(`INSERT INTO tag_list (list_id, tag_id) VALUES(?, ?)`, lastInsertIdList, lastInsertIdTag)
		if err != nil {
			_ = tx.Rollback()
			log.Fatal(err)
		}
	}

	var listDb List
	var mysqlDate string
	err = db.QueryRow(`SELECT * FROM list WHERE id = ?`, lastInsertIdList).Scan(&listDb.ID, &listDb.Title, &listDb.Body, &mysqlDate)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	rowsTags, err := db.Query(`SELECT name tag_name FROM tag_list JOIN tag ON tag.id = tag_id WHERE list_id = ?`, lastInsertIdList)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	var tags []string
	for rowsTags.Next() {
		var tag Tag
		rowsTags.Scan(&tag.Name)
		tags = append(tags, tag.Name)
	}

	listDb.Tags = tags

	t, _ := time.Parse(time.RFC3339, mysqlDate)
	listDb.CreatedAt = t.Unix()

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	defer rowsTags.Close()
	defer db.Close()

	return listDb
}