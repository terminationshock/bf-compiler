package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	template = `.text
.globl main
main:
  pushq %%rbp
  movq %%rsp, %%rbp
  movq %%rsp, %%r12
  subq $8, %%r12
  xorq %%rax, %%rax
  movq $%d, %%r13
  .loop0:
  pushq %%rax
  subq $1, %%r13
  cmpq $0, %%r13
  jne .loop0
  %sxorq %%rax, %%rax
  movq %%rbp, %%rsp
  popq %%rbp
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

func Assembly(code []*Command, file string, stackSize int) (string, error) {
	if stackSize < 1 {
		return "", errors.New("Invalid stack size")
	}

	loopId := 0
	loops := []*Loop{}

	program := ""

	for _, c := range code {
		program += fmt.Sprintf("# %s at (%d,%d)", strings.Repeat(c.String, c.Count), c.Row, c.Col) + br
		switch c.String {
		case ">":
			program += fmt.Sprintf("subq $%d, %%r12", 8 * c.Count) + br
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
			program += "movq (%r12), %rdi" + br
			program += "call putchar" + br
			break
		case ",":
			program += "call getchar" + br
			program += "movq %rax, (%r12)" + br
			break
		case "[":
			loopId++
			loops = append(loops, &Loop {
				Number: loopId,
				Row: c.Row,
				Col: c.Col,
			})
			program += fmt.Sprintf(".loop%d:", loopId) + br
			program += "cmpq $0, (%r12)" + br
			program += "je " + fmt.Sprintf(".break%d", loopId) + br
			break
		case "]":
			n := len(loops) - 1
			if n < 0 {
				return "", errors.New(fmt.Sprintf("No matching loop begin at %s:%d:%d", file, c.Row, c.Col))
			}
			loop := loops[n].Number
			mergedBreak := fmt.Sprintf(".break%d", loop)
			program += "jmp " + fmt.Sprintf(".loop%d", loop) + br
			program += mergedBreak + ":" + br
			for i := 1; i < c.Count; i++ {
				n--
				if n < 0 {
					return "", errors.New(fmt.Sprintf("No matching loop begin at %s:%d:%d", file, c.Row, c.Col))
				}
				program = strings.Replace(program, fmt.Sprintf(".break%d", loops[n].Number), mergedBreak, -1)
			}
			loops = loops[:n]
			break
		default:
			if c.MultiplyLoop != nil {
				for _, m := range c.MultiplyLoop {
					if m.Factor > 1 || m.Factor < -1 {
						program += fmt.Sprintf("imulq $%d, (%%r12), %%rax", abs(m.Factor)) + br
					} else {
						program += "movq (%r12), %rax" + br
					}
					if m.Factor > 0 {
						program += fmt.Sprintf("addq %%rax, %d(%%r12)", -8 * m.CopyTo) + br
					} else {
						program += fmt.Sprintf("subq %%rax, %d(%%r12)", -8 * m.CopyTo) + br
					}
				}
				program += "movq $0, (%r12)" + br
			}
			break
		}
	}

	if len(loops) > 0 {
		n := len(loops) - 1
		return "", errors.New(fmt.Sprintf("No matching loop end at %s:%d:%d", file, loops[n].Row, loops[n].Col))
	}

	if stackSize % 2 > 0 {
		stackSize++
	}

	return fmt.Sprintf(template, stackSize, program), nil
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
