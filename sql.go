package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetJSON() ([]byte, error) {
	return os.ReadFile(config.Json.Path)
}

func GetSQL() ([]byte, error) {
	return os.ReadFile(config.Sql.Path)
}

func (ps *Postals) Insert() (e error) {

	ps.Unique()

	db, err := NewSqlite(config.Sqlite.Path)
	if err != nil {
		e = err
		return
	}
	defer db.Close()

	db.DB.AutoMigrate(&Postal{})
	db.DB.Table(Tablename).Create(ps.List())
	ps.SaveSQL(config.Sql.Path)
	ps.SaveJSON(config.Json.Path)
	return
}

func CloseSQL(f *os.File) (e error) {
	f.Close()
	return
}

func (ps Postals) AppendSQL(f *os.File) (e error) {
	s, err := ps.BuildSQL()
	if err != nil {
		e = err
		return
	}
	fmt.Fprintln(f, s)
	return
}

func OpenSQL(dst string) (f *os.File, e error) {
	os.Remove(dst)
	return os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
}

func (ps Postals) SaveJSON(dst string) (e error) {
	j, err := json.MarshalIndent(ps, "", "  ")
	if err != nil {
		e = err
		return
	}
	if err := os.WriteFile(dst, j, 0777); err != nil {
		e = err
		return
	}
	return
}

func (ps Postals) SaveSQL(dst string) (e error) {

	s, err := ps.BuildSQL()
	if err != nil {
		e = err
		return
	}
	if err := os.WriteFile(dst, []byte(s), 0777); err != nil {
		e = err
		return
	}
	return
}

func (ps Postals) BuildSQL() (s string, e error) {
	tmpl := "INSERT INTO postal (" +
		"first, last, code, prefcode, pref, pref_kana, city, city_kana, street, street_kana, jiscode" +
		") VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');\n"
	for _, v := range ps {
		sql := fmt.Sprintf(tmpl,
			v.First,
			v.Last,
			v.Code,
			v.Prefcode,
			v.Pref,
			v.Pref_kana,
			v.City,
			v.City_kana,
			v.Street,
			v.Street_kana,
			v.Jiscode)
		s += sql
	}
	return
}
