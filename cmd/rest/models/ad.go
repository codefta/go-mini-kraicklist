package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type List struct {
	ID        int      `json:"id,omitempty"`
	Title     string   `json:"title,omitempty" validate:"required"`
	Body      string   `json:"body,omitempty" validate:"required"`
	Tags      []string `json:"tags,omitempty" validate:"required"`
	CreatedAt int64    `json:"created_at,omitempty"`
	UpdatedAt int64    `json:"updated_at,omitempty"`
}

type Tag struct {
	ListId int
	Name   string
}
