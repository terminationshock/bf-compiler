package main

import (
	"errors"
	"fmt"
)

var (
	template = `global main
extern putchar
extern getchar
extern mympi_init
extern mympi_rank
extern mympi_allreduce
extern mympi_finalize
section .text

main:
  push rbp
  mov rbp, rsp
  sub rsp, 16
  mov r12, rsp
  call mympi_init
  call mympi_rank
  mov r13, rax
  xor rax, rax
  mov [r12], rax
  %scall mympi_finalize
  mov rsp, rbp
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
	skipId := 0
	loopId := 0
	loops := []*Loop{}

	program := ""

	/*
		Special registers:
			r12: Brainfuck stack pointer
			r13: MPI rank
	*/
	for _, c := range code {
		program += fmt.Sprintf("; %c at (%d,%d)", c.Char, c.Row, c.Col) + br
		switch c.Char {
		case '>':
			skipId++
			program += fmt.Sprintf("sub r12, %d", 8 * c.Count) + br
			program += "cmp r12, rsp" + br
			program += "jge " + fmt.Sprintf(".skip%d", skipId) + br
			j := c.Count
			if j % 2 != 0 {
				j++
			}
			program += "xor rax, rax" + br
			for i := 0; i < j; i++ {
				program += "push rax" + br
			}
			program += fmt.Sprintf(".skip%d:", skipId) + br
			break
		case '<':
			program += fmt.Sprintf("add r12, %d", 8 * c.Count) + br
			break
		case '+':
			program += fmt.Sprintf("add dword [r12], %d", c.Count) + br
			break
		case '-':
			program += fmt.Sprintf("sub dword [r12], %d", c.Count) + br
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
				loopId++
				loops = append(loops, &Loop {
					Number: loopId,
					Row: c.Row,
					Col: c.Col,
				})
				program += "cmp dword [r12], 0" + br
				program += "je " + fmt.Sprintf(".break%d", loopId) + br
				program += fmt.Sprintf(".loop%d:", loopId) + br
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
		case '#':
			program += "mov [r12], r13" + br
			break
		case '$':
			for i := 0; i < c.Count; i++ {
				program += "lea rdi, [r12]" + br
				program += "call mympi_allreduce" + br
			}
			break
		}
	}

	if len(loops) > 0 {
		n := len(loops) - 1
		return "", errors.New(fmt.Sprintf("No matching loop end at %s:%d:%d", file, loops[n].Row, loops[n].Col))
	}

	return fmt.Sprintf(template, program), nil
}
