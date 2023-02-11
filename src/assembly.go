package main

import (
	"errors"
	"fmt"
)

var (
	template = `.text
.globl main
main:
  pushq %%rbp
  movq %%rsp, %%rbp
  subq $16, %%rsp
  movq %%rsp, %%r12
  call mympi_init
  call mympi_rank
  movq %%rax, %%r13
  xorq %%rax, %%rax
  movq %%rax, (%%r12)
  %scall mympi_finalize
  movq %%rbp, %%rsp
  popq %%rbp
  xorq %%rax, %%rax
  ret
.section .note.GNU-stack
`
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
		program += fmt.Sprintf("# %s at (%d,%d)", c.String, c.Row, c.Col) + br
		switch c.String {
		case ">":
			skipId++
			program += fmt.Sprintf("subq $%d, %%r12", 8 * c.Count) + br
			program += "cmpq %rsp, %r12" + br
			program += "jge " + fmt.Sprintf(".skip%d", skipId) + br
			j := c.Count
			if j % 2 != 0 {
				j++
			}
			program += "xor %rax, %rax" + br
			for i := 0; i < j; i++ {
				program += "pushq %rax" + br
			}
			program += fmt.Sprintf(".skip%d:", skipId) + br
			break
		case "<":
			program += fmt.Sprintf("addq $%d, %%r12", 8 * c.Count) + br
			break
		case "+":
			program += fmt.Sprintf("addq $%d, (%%r12)", c.Count) + br
			break
		case "-":
			program += fmt.Sprintf("subq $%d, (%%r12)", c.Count) + br
			break
		case ".":
			for i := 0; i < c.Count; i++ {
				program += "movq (%r12), %rdi" + br
				program += "call putchar" + br
			}
			break
		case ",":
			for i := 0; i < c.Count; i++ {
				program += "call getchar" + br
				program += "movq %rax, (%r12)" + br
			}
			break
		case "[":
			for i := 0; i < c.Count; i++ {
				loopId++
				loops = append(loops, &Loop {
					Number: loopId,
					Row: c.Row,
					Col: c.Col,
				})
				program += "cmpq $0, (%r12)" + br
				program += "je " + fmt.Sprintf(".break%d", loopId) + br
				program += fmt.Sprintf(".loop%d:", loopId) + br
			}
			break
		case "]":
			for i := 0; i < c.Count; i++ {
				n := len(loops) - 1
				if n < 0 {
					return "", errors.New(fmt.Sprintf("No matching loop begin at %s:%d:%d", file, c.Row, c.Col))
				}
				loop := loops[n].Number
				loops = loops[:n]
				program += "movq (%r12), %rax" + br
				program += "cmpq $0, %rax" + br
				program += "jne " + fmt.Sprintf(".loop%d", loop) + br
				program += fmt.Sprintf(".break%d:", loop) + br
			}
			break
		case "#":
			program += "movq %r13, (%r12)" + br
			break
		case "$":
			for i := 0; i < c.Count; i++ {
				program += "leaq (%r12), %rdi" + br
				program += "call mympi_allreduce" + br
			}
			break
		case SET_ZERO:
			program += "movq $0, (%r12)" + br
			break
		case ADD_LEFT:
			program += "movq (%r12), %rax" + br
			program += "addq %rax, +8(%r12)" + br
			program += "movq $0, (%r12)" + br
			break
		case ADD_RIGHT:
			skipId++
			program += "cmpq %rsp, %r12" + br
			program += "jg " + fmt.Sprintf(".skip%d", skipId) + br
			program += "xor %rax, %rax" + br
			program += "pushq %rax" + br
			program += "pushq %rax" + br
			program += fmt.Sprintf(".skip%d:", skipId) + br
			program += "movq (%r12), %rax" + br
			program += "addq %rax, -8(%r12)" + br
			program += "movq $0, (%r12)" + br
			break
		}
	}

	if len(loops) > 0 {
		n := len(loops) - 1
		return "", errors.New(fmt.Sprintf("No matching loop end at %s:%d:%d", file, loops[n].Row, loops[n].Col))
	}

	return fmt.Sprintf(template, program), nil
}
