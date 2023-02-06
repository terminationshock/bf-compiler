package main

import (
	"errors"
	"fmt"
)

var (
	template = `global main
extern putchar
extern getchar
extern MPI_Init
extern MPI_Comm_rank
extern MPI_Allreduce
extern MPI_Finalize
extern ompi_mpi_comm_world
extern ompi_mpi_int
extern ompi_mpi_op_sum
section .text

main:
  push rbp
  mov rbp, rsp
  sub rsp, 16
  mov r12, rsp
  xor rdi, rdi
  xor rsi, rsi
  call MPI_Init
  mov rdi, ompi_mpi_comm_world
  lea rsi, [r12]
  call MPI_Comm_rank
  mov r13, [r12]
  xor rax, rax
  mov [r12], rax
  mov r14, 10				
  .loop0:
  dec r14
  cmp r14, 0
  push rax
  jne .loop0
  %smov rsp, rbp
  call MPI_Finalize
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
	loopId := 0
	loops := []*Loop{}

	program := ""

	for _, c := range code {
		program += fmt.Sprintf("; %c at (%d,%d)", c.Char, c.Row, c.Col) + br
		switch c.Char {
		case '>':
			program += fmt.Sprintf("sub r12, %d", 8 * c.Count) + br
			break
		case '<':
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
				loopId++
				loops = append(loops, &Loop {
					Number: loopId,
					Row: c.Row,
					Col: c.Col,
				})
				program += "mov rax, [r12]" + br
				program += "cmp rax, 0" + br
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
			for i := 0; i < c.Count; i++ {
				program += "mov [r12], r13" + br
			}
			break
		case '$':
			for i := 0; i < c.Count; i++ {
				program += "mov rdi, 1" + br
		    	program += "lea rsi, [r12]" + br
		    	program += "mov rdx, 1" + br
				program += "mov rcx, ompi_mpi_int" + br
				program += "mov r8, ompi_mpi_op_sum" + br
				program += "mov r9, ompi_mpi_comm_world" + br
				program += "call MPI_Allreduce" + br
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
