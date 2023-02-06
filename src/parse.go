package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

type Command struct {
	Char rune
	Count int
	Row int
	Col int
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
		if strings.ContainsRune("><+-.,[]#$", r) {
			command := &Command {
				Char: rune(char),
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

	if len(code) == 0 {
		return nil, errors.New("Empty program")
	}

	cnt := 0
	optimizedCode := []*Command{code[0]}
	i := 1
	for i < len(code) {
		if optimizedCode[cnt].Char == code[i].Char {
			optimizedCode[cnt].Count++
		} else {
			optimizedCode = append(optimizedCode, code[i])
			cnt++
		}
		i++
	}

	return optimizedCode, nil
}
