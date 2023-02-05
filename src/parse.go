package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

type Command struct {
	Char rune
	Count int
}

func Parse(file string) ([]*Command, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	code := []*Command{}
	for _, char := range content {
		r := rune(char)
		if (strings.ContainsRune("><+-.,[]", r)) {
			command := &Command {
				Char: rune(char),
				Count: 1,
			}
			code = append(code, command)
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
