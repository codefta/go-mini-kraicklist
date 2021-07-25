package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

type StorageConfigs struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func NewStorage(configs StorageConfigs) (*Storage, error) {
	// initialize db object
	config := fmt.Sprintf("%s:%s@tcp(%v:%v)/%s?parseTime=true", configs.DBUser, configs.DBPass, configs.DBHost, configs.DBPort, configs.DBName)
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