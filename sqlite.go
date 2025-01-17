package main

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	SQLiteChunk = 500
)

type Sqlite struct {
	DB  *gorm.DB
	Dsn string
}

func NewSqlite(filedsn string) (d Sqlite, e error) {
	os.Remove(filedsn)

	d.DB, e = gorm.Open(sqlite.Open(filedsn), &gorm.Config{
		CreateBatchSize: 500,
	})
	if e != nil {
		return
	}
	d.Dsn = filedsn
	return
}

func (q Sqlite) Close() (e error) {
	d2, err := q.DB.DB()
	if err != nil {
		e = err
	} else {
		d2.Close()
	}
	return
}
