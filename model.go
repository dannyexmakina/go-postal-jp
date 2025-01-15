package main

import (
	"fmt"
	"os"
)

type Postal struct {
	Id          int    `json:"id" gorm:"primaryKey,column:id"`
	First       string `json:"first" gorm:"index,column:first"`
	Last        string `json:"last" gorm:"index,column:last"`
	Code        string `json:"code" gorm:"index,column:code"`
	Prefcode    string `json:"pref_code" gorm:"column:pref_code"`
	Pref        string `json:"pref" gorm:"column:pref"`
	Pref_kana   string `json:"pref_kana" gorm:"column:pref_kana"`
	City        string `json:"city" gorm:"column:city"`
	City_kana   string `json:"city_kana" gorm:"column:city_kana"`
	Street      string `json:"street" gorm:"column:street"`
	Street_kana string `json:"street_kana" gorm:"column:street_kana"`
	Jiscode     string `json:"jiscode" gorm:"index,column:jiscode"`
}
type Postals []Postal

func (p *Postal) EraseCity() *Postal {
	p.City = ""
	p.City_kana = ""
	p.Street = ""
	p.Street_kana = ""
	return p
}

func CloseSQL(f *os.File) (e error) {
	f.Close()
	return
}

func (ps Postals) AppendSQL(f *os.File) (e error) {
	s, err := ps.SQL()
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

func (ps Postals) SaveSQL(dst string) (e error) {

	s, err := ps.SQL()
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

func (ps Postals) SQL() (s string, e error) {
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
