package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB(dataSourceName string) error {
	var err error
	DB, err = sqlx.Connect("postgres", dataSourceName)
	return err
}
