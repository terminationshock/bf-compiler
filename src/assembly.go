package main

import (
	"errors"
	"fmt"
)

var (
	template = `global main
extern putchar
extern getchar
section .text
main:
  push rbp
  mov rbp, rsp
  xor rax, rax
  push rax
  mov r12, rsp
  %s%smov rsp, rbp
  pop rbp
  xor rax, rax
  ret`
	br = "\n  "
)

type Loop struct {
	Number int
	Row int
	Col int
}

func Assembly(code []*Command, file string) (string, error) {
	stack := 0
	maxStack := 0
	loopCounter := 0
	loops := []*Loop{}

	init := ""
	program := ""

	for _, c := range code {
		program += fmt.Sprintf("; %c at (%d,%d)", c.Char, c.Row, c.Col) + br
		switch c.Char {
		case '>':
			program += fmt.Sprintf("sub r12, %d", 8 * c.Count) + br
			stack += c.Count
			if stack > maxStack {
				init += "push rax" + br
				maxStack = stack
			}
			break
		case '<':
			stack -= c.Count
			if stack < 0 {
				return "", errors.New(fmt.Sprintf("Stack underflow at %s:%d:%d", file, c.Row, c.Col))
			}
			program += fmt.Sprintf("add r12, %d", 8 * c.Count) + br
			break
		case '+':
			program += "mov rax, [r12]" + br
			program += fmt.Sprintf("add rax, %d", c.Count) + br
			program += "mov [r12], rax" + br
			break
		case '-':
			program += "mov rax, [r12]" + br
			program += fmt.Sprintf("sub rax, %d", c.Count) + br
			program += "mov [r12], rax" + br
			break
		case '.':
			for i := 0; i < c.Count; i++ {
				program += "mov rdi, [r12]" + br
				program += "call putchar" + br
			}
			break
		case ',':
			for i := 0; i < c.Count; i++ {
				program += "call getchar" + br
				program += "mov [r12], rax" + br
			}
			break
		case '[':
			for i := 0; i < c.Count; i++ {
				loopCounter++
				loops = append(loops, &Loop {
					Number: loopCounter,
					Row: c.Row,
					Col: c.Col,
				})
				program += "mov rax, [r12]" + br
				program += "cmp rax, 0" + br
				program += "je " + fmt.Sprintf(".break%d", loopCounter) + br
				program += fmt.Sprintf(".loop%d:", loopCounter) + br
			}
			break
		case ']':
			for i := 0; i < c.Count; i++ {
				n := len(loops) - 1
				if n < 0 {
					return "", errors.New(fmt.Sprintf("No matching loop begin at %s:%d:%d", file, c.Row, c.Col))
				}
				loop := loops[n].Number
				loops = loops[:n]
				program += "mov rax, [r12]" + br
				program += "cmp rax, 0" + br
				program += "jne " + fmt.Sprintf(".loop%d", loop) + br
				program += fmt.Sprintf(".break%d:", loop) + br
			}
			break
		}
	}

	if len(loops) > 0 {
		n := len(loops) - 1
		return "", errors.New(fmt.Sprintf("No matching loop end at %s:%d:%d", file, loops[n].Row, loops[n].Col))
	}

	return fmt.Sprintf(template, init, program), nil
}
