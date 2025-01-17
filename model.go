package main

const (
	Tablename   = "postals"
	TempDirName = "tmp"
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
func (ps Postals) Length() (n int) {
	return len(ps.List())
}

func (ps Postals) List() (p []Postal) {
	p = ps
	return
}

func (ps *Postals) Unique() (e error) {
	var newps []Postal
	m := make(map[string]Postal)
	for _, p := range ps.List() {
		v, found := m[p.Code]
		if found {
			v.EraseCity()
			m[p.Code] = v
		} else {
			m[p.Code] = p
		}
	}
	for _, v := range m {
		newps = append(newps, v)
	}
	var np Postals
	np = newps
	ps = &np
	return
}
