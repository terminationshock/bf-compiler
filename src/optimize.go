package main

import (
	"fmt"
	"strings"
)

func Optimize(commands []*Command, verbose bool) []*Command {
	cnt := -1
	optimized := []*Command{}
	i := 0
	initialValueChanged := false
	skippedLoop := 0
	currentLoop := []*Command{}
	currentLoopBegin := 0

	if verbose {
		fmt.Println("Optimization report:")
	}
	for i < len(commands) {
		if skippedLoop > 0 {
			// Skip all commands within the first useless loop(s)
			if commands[i].String == "[" {
				skippedLoop++
			} else if commands[i].String == "]" {
				skippedLoop--
			}
			i++
		} else if !initialValueChanged && commands[i].String == "[" {
			// Whether to skip a useless loop
			skippedLoop = 1
			report(verbose, commands[i], "Skipping unreachable loop")
			i++
		} else if cnt >= 0 && optimized[cnt].String == commands[i].String && strings.Contains("><+-]", commands[i].String) {
			// Increase counter for duplicated symbols
			optimized[cnt].Count++
			i++
		} else if i < len(commands) - 1 && commands[i].String == "[" && commands[i+1].String == "]" {
			// Skip empty loop
			report(verbose, commands[i], "Skipping empty loop []")
			i += 2
		} else {
			if verbose {
				n := len(optimized) - 1
				if n >= 0 && optimized[n].Count > 1 {
					report(verbose, optimized[n], "Merging commands", getPattern(optimized[n]))
				}
			}
			optimized = append(optimized, commands[i])
			cnt++
			if !initialValueChanged {
				initialValueChanged = strings.Contains("+-,", commands[i].String)
			}
			if commands[i].String == "[" {
				currentLoop = []*Command{commands[i]}
				currentLoopBegin = cnt
			} else if len(currentLoop) > 0 {
				currentLoop = append(currentLoop, commands[i])
				if commands[i].String == "]" {
					if isMultiplyLoop(currentLoop) {
						cnt = currentLoopBegin
						optimized = optimized[:cnt+1]
						optimized[cnt].String = getLoopPattern(currentLoop)
						optimized[cnt].MultiplyLoop = newMultiplyLoop(currentLoop)
						report(verbose, currentLoop[0], "Optimizing loop", optimized[cnt].String)
					}
					currentLoop = []*Command{}
				}
			}
			i++
		}
	}

	return optimized
}

func getLoopPattern(commands []*Command) string {
	pattern := ""
	for _, c := range commands {
		pattern += getPattern(c)
	}
	return pattern
}

func getPattern(c *Command) string {
	return strings.Repeat(c.String, c.Count)
}

func isMultiplyLoop(commands []*Command) bool {
	pointer := 0
	found := false
	for _, c := range commands {
		if c.String == "<" {
			pointer -= c.Count
		} else if c.String == ">" {
			pointer += c.Count
		} else if c.String == "+" {
			if pointer == 0 {
				return false
			}
		} else if c.String == "-" {
			if pointer == 0 {
				if c.Count > 1 || found {
					return false
				}
				found = true
			}
		} else if c.String == "." || c.String == "," {
			return false
		}
	}
	return found && pointer == 0
}

func newMultiplyLoop(commands []*Command) []*Multiply {
	m := []*Multiply{}

	pointer := 0
	for _, c := range commands {
		switch c.String {
		case "<":
			pointer -= c.Count
			break
		case ">":
			pointer += c.Count
			break
		case "+":
			m = append(m, &Multiply{
				CopyTo: pointer,
				Factor: c.Count,
			})
			break
		case "-":
			if pointer != 0 {
				m = append(m, &Multiply{
					CopyTo: pointer,
					Factor: -c.Count,
				})
			}
			break
		}
	}
	return m
}

func report(verbose bool, command *Command, message ...string) {
	if verbose {
		pos := ""
		if command.Row < 10 {
			pos += " "
		}
		pos = fmt.Sprintf("%s(%d:%d)", pos, command.Row, command.Col)
		fill := strings.Repeat(" ", 9 - len(pos))
		fmt.Printf("  %s %s%s\n", pos, fill, strings.Join(message, " "))
	}
}
