package config

import (
	"database/sql"
	"final-project-go/utilitis"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	username = utilitis.Getenv("DATABASE_USERNAME", "root")
	password = utilitis.Getenv("DATABASE_PASSWORD", "")
	host     = utilitis.Getenv("DATABASE_HOST", "127.0.0.1")
	port     = utilitis.Getenv("DATABASE_PORT", "3306")
	database = utilitis.Getenv("DATABASE_NAME", "project_go")
)

var (
	dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
)

func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
