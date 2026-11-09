package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

const (
	AvoidAddress1 = "以下に掲載がない場合"
	AvoidAddress2 = "イカニケイサイガナイバアイ"
)

func ParseFile(filename string) (ps Postals, e error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		e = err
		return
	}
	b = bytes.TrimSpace(b)
	lines := strings.Split(strings.ReplaceAll(string(b), "\r\n", "\n"), "\n")
	for _, line := range lines {
		p, err := ParseLine(line)
		if err != nil {
			log.Printf("parse error: %s, line:%s", err.Error(), line)
			continue
		}
		ps = append(ps, p)
	}

	return
}

func ParseLine(line string) (p Postal, e error) {
	var r1 = regexp.MustCompile("（(.*?)）")
	var r2 = regexp.MustCompile("\\((.*?)\\)")
	line = strings.Replace(line, "\"", "", -1)

	line = KanaHalfToFull(line)
	if len(line) < 1 {
		e = errors.New("line empty")
		return
	}
	cells := strings.Split(line, ",")
	if len(cells) < 8 {
		e = errors.Errorf("short line: %s", line)
		return
	}

	code := cells[2]
	jiscode := cells[0]
	street := r1.ReplaceAllString(cells[8], "")
	street_kana := r2.ReplaceAllString(cells[5], "")

	if street == AvoidAddress1 || street_kana == AvoidAddress2 {
		street = ""
		street_kana = ""
	}

	p = Postal{
		First:       code[0:3],
		Last:        code[3:],
		Code:        code,
		Prefcode:    jiscode[0:2],
		Pref:        cells[6],
		Pref_kana:   cells[3],
		City:        cells[7],
		City_kana:   cells[4],
		Street:      street,
		Street_kana: street_kana,
		Jiscode:     jiscode,
	}
	return
}
