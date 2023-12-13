package storage

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

const (
	dns = "usi1p2fhaqjrbcxd:7b3gioYphTLYaZoRbF8P@tcp(byr2cqsjyxldxj3oabah-mysql.services.clever-cloud.com:3306)/byr2cqsjyxldxj3oabah"
)

func NewConnection() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", dns)
		if err != nil {
			log.Fatalf("Can't open database %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping %v", err)
		}
	})
}

func Pool() *sql.DB {
	return db
}
