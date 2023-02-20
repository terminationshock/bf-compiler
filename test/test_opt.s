.text
.globl main
main:
  pushq %rbp
  movq %rsp, %rbp
  movq %rsp, %r12
  subq $8, %r12
  subq $240000, %rsp
  xorq %rax, %rax
  movq $30000, %rcx
  leaq (%rsp), %rdi
  rep stosq
  # > at (2,1)
  subq $8, %r12
  # ++++ at (4,1)
  addq $4, (%r12)
  # [ at (5,1)
  .loop1:
  cmpq $0, (%r12)
  je .break1
  # > at (5,2)
  subq $8, %r12
  # ] at (5,3)
  jmp .loop1
  .break1:
  # >>> at (7,1)
  subq $24, %r12
  # +++ at (9,1)
  addq $3, (%r12)
  # [>>->+<+++<<-] at (10,1)
  movq (%r12), %rax
  subq %rax, -16(%r12)
  movq (%r12), %rax
  addq %rax, -24(%r12)
  imulq $3, (%r12), %rax
  addq %rax, -16(%r12)
  movq $0, (%r12)
  # + at (11,1)
  addq $1, (%r12)
  # [->-<] at (12,1)
  movq (%r12), %rax
  subq %rax, -8(%r12)
  movq $0, (%r12)
  # [-] at (14,1)
  movq $0, (%r12)
  # + at (15,1)
  addq $1, (%r12)
  # [ at (16,1)
  .loop2:
  cmpq $0, (%r12)
  je .break4
  # + at (16,2)
  addq $1, (%r12)
  # [ at (16,3)
  .loop3:
  cmpq $0, (%r12)
  je .break4
  # + at (16,4)
  addq $1, (%r12)
  # [ at (16,5)
  .loop4:
  cmpq $0, (%r12)
  je .break4
  # + at (16,6)
  addq $1, (%r12)
  # ]]] at (16,7)
  jmp .loop4
  .break4:
  # + at (17,1)
  addq $1, (%r12)
  # [-] at (18,3)
  movq $0, -16(%r12)
  # ++ at (18,7)
  addq $2, -24(%r12)
  # ++ at (18,10)
  addq $2, -32(%r12)
  # - at (18,13)
  subq $1, -24(%r12)
  # >>>> at (17,1)
  subq $32, %r12
  # [ at (19,1)
  .loop5:
  cmpq $0, (%r12)
  je .break5
  # + at (19,2)
  addq $1, (%r12)
  # ] at (19,3)
  jmp .loop5
  .break5:
  # . at (20,3)
  movq -16(%r12), %rdi
  call putchar
  # - at (20,6)
  subq $1, (%r12)
  # [-] at (20,8)
  movq $0, -8(%r12)
  # >>>> at (20,1)
  subq $32, %r12
  # + at (20,14)
  addq $1, (%r12)
  xorq %rax, %rax
  movq %rbp, %rsp
  popq %rbp
  ret
.section .note.GNU-stack
