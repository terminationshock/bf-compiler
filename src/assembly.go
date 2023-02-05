package main

import (
	"errors"
	"fmt"
)

var (
	template = `global main
extern putchar
section .text
main:
  push  rbp
  mov   rbp, rsp
  xor   rax, rax
  push  rax
  mov	r12, rsp
  %s
  %s
  mov   rsp, rbp
  pop   rbp
  xor   rax, rax
  ret`
	br = "\n"
)

func Assembly(code []*Command) (string, error) {
	stack := 0
	maxStack := 0
	loop := 0

	init := ""
	program := ""

	for _, c := range code {
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
				return "", errors.New("Stack underflow")
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
			return "", errors.New(", not implemented")
			break
		case '[':
			for i := 0; i < c.Count; i++ {
				loop++
				program += "mov rax, [r12]" + br
				program += "cmp rax, 0" + br
				program += "je " + fmt.Sprintf(".break%d", loop) + br
				program += fmt.Sprintf(".loop%d:", loop) + br
			}
			break
		case ']':
			for i := 0; i < c.Count; i++ {
				program += "mov rax, [r12]" + br
				program += "cmp rax, 0" + br
				program += "jne " + fmt.Sprintf(".loop%d", loop) + br
				program += fmt.Sprintf(".break%d:", loop) + br
				loop--
				if loop < 0 {
					return "", errors.New("No matching loop begin")
				}
			}
			break
		}
	}

	return fmt.Sprintf(template, init, program), nil
}
