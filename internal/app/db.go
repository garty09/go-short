package app

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB(datasource string) *sql.DB {
	var db *sql.DB
	var err error
	db, err = sql.Open(datasource, os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp(mysql:"+os.Getenv("DB_PORT")+")/"+os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Println(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db
}
