package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

type Command struct {
	String string
	Count int
	Row int
	Col int
}

func Parse(file string) ([]*Command, bool, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, false, err
	}

	hasMpi := false
	row := 1
	col := 1
	code := []*Command{}
	for _, char := range content {
		r := rune(char)
		if strings.ContainsRune("><+-.,[]#$?", r) {
			command := &Command {
				String: string(char),
				Count: 1,
				Row: row,
				Col: col,
			}
			code = append(code, command)
			if strings.ContainsRune("#$", r) {
				hasMpi = true
			}
		}
		if char == '\n' {
			row++
			col = 1
		} else {
			col++
		}
	}

	if len(code) == 0 {
		return nil, false, errors.New("Empty program")
	}

	return code, hasMpi, nil
}
