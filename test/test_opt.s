.text
.globl main
main:
  pushq %rbp
  movq %rsp, %rbp
  movq %rsp, %r12
  subq $8, %r12
  xorq %rax, %rax
  movq $30000, %r13
  .loop0:
  pushq %rax
  subq $1, %r13
  testq %r13, %r13
  jne .loop0
  subq $8, %r12
  addq $4, (%r12)
  .loop1:
  cmpq $0, (%r12)
  je .break1
  subq $8, %r12
  jmp .loop1
  .break1:
  subq $24, %r12
  addq $3, (%r12)
  movq (%r12), %rax
  subq %rax, -16(%r12)
  movq (%r12), %rax
  addq %rax, -24(%r12)
  imulq $3, (%r12), %rax
  addq %rax, -16(%r12)
  movq $0, (%r12)
  addq $1, (%r12)
  movq (%r12), %rax
  subq %rax, -8(%r12)
  movq $0, (%r12)
  addq $4, (%r12)
  movq $0, (%r12)
  addq $1, (%r12)
  .loop2:
  cmpq $0, (%r12)
  je .break4
  addq $1, (%r12)
  .loop3:
  cmpq $0, (%r12)
  je .break4
  addq $1, (%r12)
  .loop4:
  cmpq $0, (%r12)
  je .break4
  addq $1, (%r12)
  jmp .loop4
  .break4:
  addq $1, (%r12)
  movq $0, -16(%r12)
  addq $2, -24(%r12)
  addq $2, -32(%r12)
  subq $1, -24(%r12)
  subq $32, %r12
  .loop5:
  cmpq $0, (%r12)
  je .break5
  addq $1, (%r12)
  jmp .loop5
  .break5:
  movq -16(%r12), %rdi
  call putchar
  subq $1, (%r12)
  movq $0, -8(%r12)
  subq $32, %r12
  addq $1, (%r12)
  xorq %rax, %rax
  movq %rbp, %rsp
  popq %rbp
  ret
.section .note.GNU-stack
