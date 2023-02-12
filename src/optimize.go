package main

import (
	"strings"
)

const (
	EMPTY_LOOP = "[]"
	SET_ZERO = "[-]"
	ADD_LEFT = "[<+>-]"
	ADD_RIGHT = "[>+<-]"
)

func Optimize(commands []*Command) []*Command {
	cnt := -1
	optimized := []*Command{}
	i := 0
	loop := 0
	initialValueChanged := false
	for i < len(commands) {
		if loop > 0 {
			// Skip all commands within the first useless loop(s)
			if commands[i].String == "[" {
				loop++
			} else if commands[i].String == "]" {
				loop--
			}
			i++
		} else if cnt >= 0 && optimized[cnt].String == commands[i].String && strings.Contains("><+-", commands[i].String) {
			// Increase counter for duplicated symbols
			optimized[cnt].Count++
			i++
		} else if !initialValueChanged && commands[i].String == "[" {
			// Whether to skip a useless loop
			loop = 1
			i++
		} else if isPattern(EMPTY_LOOP, i, commands) {
			optimized = append(optimized, newCommand(EMPTY_LOOP, commands[i]))
			cnt++
			i += 2
		} else if isPattern(SET_ZERO, i, commands) {
			optimized = append(optimized, newCommand(SET_ZERO, commands[i]))
			cnt++
			i += 3
		} else if isPattern(ADD_LEFT, i, commands) {
			optimized = append(optimized, newCommand(ADD_LEFT, commands[i]))
			cnt++
			i += 6
		} else if isPattern(ADD_RIGHT, i, commands) {
			optimized = append(optimized, newCommand(ADD_RIGHT, commands[i]))
			cnt++
			i += 6
		} else {
			optimized = append(optimized, commands[i])
			if !initialValueChanged {
				initialValueChanged = strings.Contains("+-,", commands[i].String)
			}
			cnt++
			i++
		}
	}

	return optimized
}

func isPattern(pattern string, i int, commands []*Command) bool {
	n := len(commands)
	if i > n - len(pattern) {
		return false
	}

	for j := 0; j < len(pattern); j++ {
		if commands[i+j].String != string(pattern[j]) {
			return false
		}
	}

	return true
}

func newCommand(pattern string, command *Command) *Command {
	return &Command {
		String: pattern,
		Count: 1,
		Row: command.Row,
		Col: command.Col,
	}
}
