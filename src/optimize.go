package main

import (
	"fmt"
	"strings"
)

func Optimize(commands []*Command, verbose bool) []*Command {
	if verbose {
		fmt.Println("Optimization report:")
	}

	optimized := optimizeUselessLoops(commands, verbose)
	optimized = optimizeDuplicatedCommands(optimized, "><+-", verbose)
	optimized = optimizeMultiplyLoops(optimized, verbose)
	optimized = optimizeDuplicatedCommands(optimized, "]", verbose)
	optimized = optimizeUselessValueChanges(optimized, verbose)
	optimized = optimizePointerMovement(optimized, verbose)

	return optimized
}

func optimizeUselessLoops(commands []*Command, verbose bool) []*Command {
	optimized := []*Command{}
	i := 0
	initialValueChanged := false
	skippedLoop := 0
	for i < len(commands) {
		if skippedLoop > 0 {
			if commands[i].String == "[" {
				skippedLoop++
			} else if commands[i].String == "]" {
				skippedLoop--
			}
			i++
		} else if !initialValueChanged && commands[i].String == "[" {
			skippedLoop = 1
			report(verbose, commands[i], "Skipping unreachable loop")
			i++
		} else if i > 0 && commands[i-1].String == "]" && commands[i].String == "[" {
			skippedLoop = 1
			report(verbose, commands[i], "Skipping unreachable loop")
			i++
		} else if i < len(commands) - 1 && commands[i].String == "[" && commands[i+1].String == "]" {
			report(verbose, commands[i], "Skipping empty loop []")
			i += 2
		} else {
			optimized = append(optimized, commands[i])
			if !initialValueChanged {
				initialValueChanged = strings.Contains("+-,", commands[i].String)
			}
			i++
		}
	}
	return optimized
}

func optimizeDuplicatedCommands(commands []*Command, pattern string, verbose bool) []*Command {
	optimized := []*Command{}
	cnt := -1
	i := 0
	for i < len(commands) {
		if cnt >= 0 && optimized[cnt].String == commands[i].String && strings.Contains(pattern, commands[i].String) {
			optimized[cnt].Count++
		} else {
			if verbose {
				n := len(optimized) - 1
				if n >= 0 && optimized[n].Count > 1 && strings.Contains(pattern, optimized[n].String) {
					report(verbose, optimized[n], "Merging commands", getPattern(optimized[n]))
				}
			}
			optimized = append(optimized, commands[i])
			cnt++
		}
		i++
	}
	return optimized
}

func optimizeMultiplyLoops(commands []*Command, verbose bool) []*Command {
	optimized := []*Command{}
	cnt := -1
	i := 0
	currentLoop := []*Command{}
	currentLoopBegin := 0
	for i < len(commands) {
		optimized = append(optimized, commands[i])
		cnt++
		if commands[i].String == "[" {
			currentLoop = []*Command{commands[i]}
			currentLoopBegin = cnt
		} else if len(currentLoop) > 0 {
			currentLoop = append(currentLoop, commands[i])
			if commands[i].String == "]" {
				if isMultiplyLoop(currentLoop) {
					cnt = currentLoopBegin
					optimized = optimized[:cnt+1]
					optimized[cnt].String = getBlockPattern(currentLoop)
					optimized[cnt].MultiplyLoop = getMultiplyLoop(currentLoop)
					report(verbose, currentLoop[0], "Optimizing loop", optimized[cnt].String)
				}
				currentLoop = []*Command{}
			}
		}
		i++
	}
	return optimized
}

func optimizeUselessValueChanges(commands []*Command, verbose bool) []*Command {
	optimized := []*Command{}
	i := 0
	block := []*Command{}
	for i < len(commands) {
		if strings.Contains("+-", commands[i].String) && i < len(commands) - 1 {
			block = append(block, commands[i])
		} else {
			if len(block) > 0 && (commands[i].String == "," || commands[i].String == "[-]") {
				report(verbose, block[0], "Skipping overwritten value", getBlockPattern(block))
				block = []*Command{}
			}
			optimized = append(optimized, block...)
			optimized = append(optimized, commands[i])
			block = []*Command{}
		}
		i++
	}
	return optimized
}

func optimizePointerMovement(commands []*Command, verbose bool) []*Command {
	optimized := []*Command{}
	i := 0
	block := []*Command{}
	for i < len(commands) {
		if (strings.Contains("><+-,.", commands[i].String) || commands[i].String == "[-]") && i < len(commands) - 1 {
			block = append(block, commands[i])
		} else {
			if len(block) > 0 && isUnoptimizedPointerMovement(block) {
				report(verbose, block[0], "Optimizing pointer movement", getBlockPattern(block))
				optimized = append(optimized, getRemoteCommands(block)...)
				pointer, str := getNetPointerMovement(block)
				if pointer != 0 {
					optimized = append(optimized, &Command {
						String: str,
						Count: pointer,
						Offset: 0,
						Row: block[0].Row,
						Col: block[0].Col,
					})
				}
				block = []*Command{}
			}
			optimized = append(optimized, block...)
			optimized = append(optimized, commands[i])
			block = []*Command{}
		}
		i++
	}
	return optimized
}

func getBlockPattern(commands []*Command) string {
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

func isUnoptimizedPointerMovement(commands []*Command) bool {
	numMove := 0
	for _, c := range commands {
		if c.String == ">" || c.String == "<" {
			numMove++
		}
	}
	return numMove > 1
}

func getNetPointerMovement(commands []*Command) (int, string) {
	pointer := 0
	for _, c := range commands {
		if c.String == ">" {
			pointer += c.Count
		} else if c.String == "<" {
			pointer -= c.Count
		}
	}
	if pointer > 0 {
		return pointer, ">"
	} else if pointer < 0 {
		return -pointer, "<"
	}
	return 0, ""
}

func getRemoteCommands(commands []*Command) []*Command {
	remote := []*Command{}
	pointer := 0
	for _, c := range commands {
		if c.String == "<" {
			pointer -= c.Count
		} else if c.String == ">" {
			pointer += c.Count
		} else {
			remote = append(remote, &Command {
				String: c.String,
				Count: c.Count,
				Offset: pointer,
				Row: c.Row,
				Col: c.Col,
				MultiplyLoop: c.MultiplyLoop,
			})
		}
	}
	return remote
}

func getMultiplyLoop(commands []*Command) []*Multiply {
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
				Offset: pointer,
				Factor: c.Count,
			})
			break
		case "-":
			if pointer != 0 {
				m = append(m, &Multiply{
					Offset: pointer,
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
