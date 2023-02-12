package main

import (
	"io/ioutil"
	"strings"
)

type Command struct {
	String string
	Count int
	Row int
	Col int
	MultiplyLoop []*Multiply
}

type Multiply struct {
	CopyTo int
	Factor int
}

func Parse(file string) ([]*Command, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	row := 1
	col := 1
	code := []*Command{}
	for _, char := range content {
		r := rune(char)
		if strings.ContainsRune("><+-.,[]", r) {
			command := &Command {
				String: string(char),
				Count: 1,
				Row: row,
				Col: col,
			}
			code = append(code, command)
		}
		if char == '\n' {
			row++
			col = 1
		} else {
			col++
		}
	}

	return code, nil
}
