package main

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	Memory *gorm.DB
	File   *gorm.DB
	Dsn    SqliteDsn
}

type SqliteDsn struct {
	Memory string
	File   string
}

func NewSqlite(memdsn, filedsn string) (d Sqlite, e error) {
	os.Remove(filedsn)
	d.Memory, e = gorm.Open(sqlite.Open(memdsn), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if e != nil {
		return
	}
	d.File, e = gorm.Open(sqlite.Open(filedsn), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if e != nil {
		return
	}
	d.Dsn.Memory = memdsn
	d.Dsn.File = filedsn
	return
}

func (q Sqlite) Save(table string) (e error) {
	e = q.File.Exec(fmt.Sprintf("ATTACH DATABASE '%s' AS mem", q.Dsn.Memory)).Error
	if e != nil {
		return
	}
	e = q.File.Exec(fmt.Sprintf("CREATE TABLE main.%s AS SELECT * FROM mem.%s", table, table)).Error
	if e != nil {
		return
	}
	e = q.File.Exec("DETACH DATABASE mem").Error
	return
}

func (q Sqlite) Close() (e error) {
	d1, err := q.Memory.DB()
	if err != nil {
		e = err
	} else {
		d1.Close()
	}
	d2, err := q.File.DB()
	if err != nil {
		e = err
	} else {
		d2.Close()
	}
	return
}
