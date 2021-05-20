package main

import (
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
