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
  # > at (5,1)
  subq $8, %r12
  # + at (5,2)
  addq $1, (%r12)
  # [ at (5,3)
  .loop1:
  cmpq $0, (%r12)
  je .break1
  # - at (5,4)
  subq $1, (%r12)
  # > at (5,5)
  subq $8, %r12
  # , at (5,6)
  call getchar
  movq %rax, (%r12)
  # ---------- at (5,7)
  subq $10, (%r12)
  # [ at (5,17)
  .loop2:
  cmpq $0, (%r12)
  je .break2
  # + at (5,19)
  addq $1, 8(%r12)
  # -------------------------------------- at (5,21)
  subq $38, (%r12)
  # > at (5,18)
  subq $8, %r12
  # [>+>+<<-] at (5,60)
  movq (%r12), %rax
  addq %rax, -8(%r12)
  addq %rax, -16(%r12)
  movq $0, (%r12)
  # >> at (5,69)
  subq $16, %r12
  # [<<+>>-] at (5,71)
  movq (%r12), %rax
  addq %rax, 16(%r12)
  movq $0, (%r12)
  # +++++++++ at (6,10)
  addq $9, -16(%r12)
  # >> at (6,6)
  subq $16, %r12
  # [ at (6,19)
  .loop3:
  cmpq $0, (%r12)
  je .break3
  # <<< at (6,20)
  addq $24, %r12
  # [>+>+<<-] at (6,23)
  movq (%r12), %rax
  addq %rax, -8(%r12)
  addq %rax, -16(%r12)
  movq $0, (%r12)
  # >> at (6,32)
  subq $16, %r12
  # [<<+>>-] at (6,34)
  movq (%r12), %rax
  addq %rax, 16(%r12)
  movq $0, (%r12)
  # < at (6,42)
  addq $8, %r12
  # [<<+>>-] at (6,43)
  movq (%r12), %rax
  addq %rax, 16(%r12)
  movq $0, (%r12)
  # >> at (6,51)
  subq $16, %r12
  # - at (6,53)
  subq $1, (%r12)
  # ] at (6,54)
  jmp .loop3
  .break3:
  # [-] at (6,58)
  movq $0, 24(%r12)
  # <<<<< at (6,55)
  addq $40, %r12
  # [>+<-] at (6,63)
  movq (%r12), %rax
  addq %rax, -8(%r12)
  movq $0, (%r12)
  # ] at (6,69)
  jmp .loop2
  .break2:
  # < at (6,70)
  addq $8, %r12
  # ] at (6,71)
  jmp .loop1
  .break1:
  # >> at (6,72)
  subq $16, %r12
  # [<<+>>-] at (7,2)
  movq (%r12), %rax
  addq %rax, 16(%r12)
  movq $0, (%r12)
  # [-] at (9,2)
  movq $0, 24(%r12)
  # << at (7,10)
  addq $16, %r12
  # [<+>>>>>>>>++++++++++<<<<<<<-] at (29,1)
  movq (%r12), %rax
  addq %rax, 8(%r12)
  imulq $10, (%r12), %rax
  addq %rax, -56(%r12)
  movq $0, (%r12)
  # > at (29,31)
  subq $8, %r12
  # +++++ at (29,32)
  addq $5, (%r12)
  # [<+++++++++>-] at (29,37)
  imulq $9, (%r12), %rax
  addq %rax, 8(%r12)
  movq $0, (%r12)
  # + at (29,51)
  addq $1, (%r12)
  # >>>>>> at (29,52)
  subq $48, %r12
  # + at (29,58)
  addq $1, (%r12)
  # [ at (29,59)
  .loop4:
  cmpq $0, (%r12)
  je .break4
  # << at (29,60)
  addq $16, %r12
  # +++ at (29,62)
  addq $3, (%r12)
  # [ at (29,65)
  .loop5:
  cmpq $0, (%r12)
  je .break5
  # >> at (29,66)
  subq $16, %r12
  # [ at (29,68)
  .loop6:
  cmpq $0, (%r12)
  je .break6
  # - at (29,69)
  subq $1, (%r12)
  # < at (29,70)
  addq $8, %r12
  # ] at (29,71)
  jmp .loop6
  .break6:
  # < at (29,72)
  addq $8, %r12
  # [ at (29,73)
  .loop7:
  cmpq $0, (%r12)
  je .break7
  # > at (29,74)
  subq $8, %r12
  # ] at (29,75)
  jmp .loop7
  .break7:
  # < at (29,76)
  addq $8, %r12
  # - at (29,77)
  subq $1, (%r12)
  # ] at (29,78)
  jmp .loop5
  .break5:
  # >> at (29,79)
  subq $16, %r12
  # [ at (30,1)
  .loop8:
  cmpq $0, (%r12)
  je .break8
  # + at (30,3)
  addq $1, -8(%r12)
  # >> at (30,2)
  subq $16, %r12
  # ] at (30,5)
  jmp .loop8
  .break8:
  # < at (30,6)
  addq $8, %r12
  # [ at (30,7)
  .loop9:
  cmpq $0, (%r12)
  je .break9
  # < at (30,8)
  addq $8, %r12
  # ] at (30,9)
  jmp .loop9
  .break9:
  # > at (30,10)
  subq $8, %r12
  # ] at (30,11)
  jmp .loop4
  .break4:
  # > at (30,12)
  subq $8, %r12
  # [ at (30,13)
  .loop10:
  cmpq $0, (%r12)
  je .break10
  # [->>>>+<<<<] at (30,14)
  movq (%r12), %rax
  addq %rax, -32(%r12)
  movq $0, (%r12)
  # +++ at (30,29)
  addq $3, -24(%r12)
  # - at (30,33)
  subq $1, -32(%r12)
  # >>>> at (30,26)
  subq $32, %r12
  # ] at (30,34)
  jmp .loop10
  .break10:
  # < at (30,35)
  addq $8, %r12
  # [ at (30,36)
  .loop11:
  cmpq $0, (%r12)
  je .break11
  # <<<< at (30,37)
  addq $32, %r12
  # ] at (30,41)
  jmp .loop11
  .break11:
  # <<<<<<<< at (30,42)
  addq $64, %r12
  # + at (30,50)
  addq $1, (%r12)
  # [ at (30,51)
  .loop12:
  cmpq $0, (%r12)
  je .break12
  # - at (30,52)
  subq $1, (%r12)
  # >>>>>>>>>>>> at (30,53)
  subq $96, %r12
  # [ at (30,65)
  .loop13:
  cmpq $0, (%r12)
  je .break13
  # < at (30,66)
  addq $8, %r12
  # + at (30,67)
  addq $1, (%r12)
  # [->>>>+<<<<] at (30,68)
  movq (%r12), %rax
  addq %rax, -32(%r12)
  movq $0, (%r12)
  # >>>>> at (30,80)
  subq $40, %r12
  # ] at (31,5)
  jmp .loop13
  .break13:
  # <<<< at (31,6)
  addq $32, %r12
  # [ at (31,10)
  .loop14:
  cmpq $0, (%r12)
  je .break14
  # >>>>> at (31,11)
  subq $40, %r12
  # [<<<<+>>>>-] at (31,16)
  movq (%r12), %rax
  addq %rax, 32(%r12)
  movq $0, (%r12)
  # <<<<< at (31,28)
  addq $40, %r12
  # - at (31,33)
  subq $1, (%r12)
  # [<<++++++++++>>-] at (31,34)
  imulq $10, (%r12), %rax
  addq %rax, 16(%r12)
  movq $0, (%r12)
  # >>> at (31,51)
  subq $24, %r12
  # [ at (31,54)
  .loop15:
  cmpq $0, (%r12)
  je .break15
  # << at (31,55)
  addq $16, %r12
  # [<+<<+>>>-] at (31,57)
  movq (%r12), %rax
  addq %rax, 8(%r12)
  addq %rax, 24(%r12)
  movq $0, (%r12)
  # < at (31,68)
  addq $8, %r12
  # [>+<-] at (31,69)
  movq (%r12), %rax
  addq %rax, -8(%r12)
  movq $0, (%r12)
  # ++ at (31,76)
  addq $2, 8(%r12)
  # + at (31,80)
  addq $1, 24(%r12)
  # - at (32,7)
  subq $1, -24(%r12)
  # >>> at (31,75)
  subq $24, %r12
  # ] at (32,8)
  jmp .loop15
  .break15:
  # [-] at (32,11)
  movq $0, 16(%r12)
  # - at (32,16)
  subq $1, 32(%r12)
  # <<<<< at (32,9)
  addq $40, %r12
  # [ at (32,18)
  .loop16:
  cmpq $0, (%r12)
  je .break16
  # - at (32,19)
  subq $1, (%r12)
  # + at (32,22)
  addq $1, -16(%r12)
  # - at (32,24)
  subq $1, -8(%r12)
  # > at (32,19)
  subq $8, %r12
  # [ at (32,25)
  .loop17:
  cmpq $0, (%r12)
  je .break17
  # >>> at (32,26)
  subq $24, %r12
  # ] at (32,29)
  jmp .loop17
  .break17:
  # > at (32,30)
  subq $8, %r12
  # [ at (32,31)
  .loop18:
  cmpq $0, (%r12)
  je .break18
  # [<+>-] at (32,32)
  movq (%r12), %rax
  addq %rax, 8(%r12)
  movq $0, (%r12)
  # + at (32,39)
  addq $1, -8(%r12)
  # >>> at (32,38)
  subq $24, %r12
  # ] at (32,42)
  jmp .loop18
  .break18:
  # <<<<< at (32,43)
  addq $40, %r12
  # ] at (32,48)
  jmp .loop16
  .break16:
  # [-] at (32,50)
  movq $0, -8(%r12)
  # + at (32,54)
  addq $1, -16(%r12)
  # - at (32,58)
  subq $1, 8(%r12)
  # < at (32,49)
  addq $8, %r12
  # [>>+<<-] at (32,59)
  movq (%r12), %rax
  addq %rax, -16(%r12)
  movq $0, (%r12)
  # < at (32,67)
  addq $8, %r12
  # ] at (32,68)
  jmp .loop14
  .break14:
  # + at (32,73)
  addq $1, 32(%r12)
  # [-] at (33,2)
  movq $0, -32(%r12)
  # >>>>> at (32,69)
  subq $40, %r12
  # [<<<+>>>-] at (33,6)
  movq (%r12), %rax
  addq %rax, 24(%r12)
  movq $0, (%r12)
  # ++++++++++ at (33,18)
  addq $10, 16(%r12)
  # <<< at (33,16)
  addq $24, %r12
  # [ at (33,29)
  .loop19:
  cmpq $0, (%r12)
  je .break19
  # - at (33,30)
  subq $1, (%r12)
  # + at (33,33)
  addq $1, -16(%r12)
  # - at (33,35)
  subq $1, -8(%r12)
  # > at (33,30)
  subq $8, %r12
  # [ at (33,36)
  .loop20:
  cmpq $0, (%r12)
  je .break20
  # >>> at (33,37)
  subq $24, %r12
  # ] at (33,40)
  jmp .loop20
  .break20:
  # > at (33,41)
  subq $8, %r12
  # [ at (33,42)
  .loop21:
  cmpq $0, (%r12)
  je .break21
  # [<+>-] at (33,43)
  movq (%r12), %rax
  addq %rax, 8(%r12)
  movq $0, (%r12)
  # + at (33,50)
  addq $1, -8(%r12)
  # >>> at (33,49)
  subq $24, %r12
  # ] at (33,53)
  jmp .loop21
  .break21:
  # <<<<< at (33,54)
  addq $40, %r12
  # ] at (33,59)
  jmp .loop19
  .break19:
  # [-] at (33,61)
  movq $0, -8(%r12)
  # + at (33,65)
  addq $1, -16(%r12)
  # >>> at (33,60)
  subq $24, %r12
  # [<<+<+>>>-] at (33,67)
  movq (%r12), %rax
  addq %rax, 16(%r12)
  addq %rax, 24(%r12)
  movq $0, (%r12)
  # + at (34,2)
  addq $1, 32(%r12)
  # + at (34,4)
  addq $1, 40(%r12)
  # <<< at (33,78)
  addq $24, %r12
  # [ at (34,7)
  .loop22:
  cmpq $0, (%r12)
  je .break30
  # - at (34,8)
  subq $1, (%r12)
  # [ at (34,9)
  .loop23:
  cmpq $0, (%r12)
  je .break30
  # - at (34,10)
  subq $1, (%r12)
  # [ at (34,11)
  .loop24:
  cmpq $0, (%r12)
  je .break30
  # - at (34,12)
  subq $1, (%r12)
  # [ at (34,13)
  .loop25:
  cmpq $0, (%r12)
  je .break30
  # - at (34,14)
  subq $1, (%r12)
  # [ at (34,15)
  .loop26:
  cmpq $0, (%r12)
  je .break30
  # - at (34,16)
  subq $1, (%r12)
  # [ at (34,17)
  .loop27:
  cmpq $0, (%r12)
  je .break30
  # - at (34,18)
  subq $1, (%r12)
  # [ at (34,19)
  .loop28:
  cmpq $0, (%r12)
  je .break30
  # - at (34,20)
  subq $1, (%r12)
  # [ at (34,21)
  .loop29:
  cmpq $0, (%r12)
  je .break30
  # - at (34,22)
  subq $1, (%r12)
  # [ at (34,23)
  .loop30:
  cmpq $0, (%r12)
  je .break30
  # - at (34,24)
  subq $1, (%r12)
  # - at (34,26)
  subq $1, 8(%r12)
  # [-<+<->>] at (34,28)
  movq (%r12), %rax
  addq %rax, 8(%r12)
  movq (%r12), %rax
  subq %rax, 16(%r12)
  movq $0, (%r12)
  # ]]]]]]]]] at (34,37)
  jmp .loop30
  .break30:
  # < at (34,46)
  addq $8, %r12
  # [ at (34,47)
  .loop31:
  cmpq $0, (%r12)
  je .break35
  # +++++ at (34,48)
  addq $5, (%r12)
  # [<<<++++++++<++++++++>>>>-] at (34,53)
  imulq $8, (%r12), %rax
  addq %rax, 24(%r12)
  addq %rax, 32(%r12)
  movq $0, (%r12)
  # + at (35,4)
  addq $1, 32(%r12)
  # - at (35,6)
  subq $1, 40(%r12)
  # < at (34,80)
  addq $8, %r12
  # [>+<<<+++++++++<->>>-] at (35,11)
  movq (%r12), %rax
  addq %rax, -8(%r12)
  imulq $9, (%r12), %rax
  addq %rax, 16(%r12)
  movq (%r12), %rax
  subq %rax, 24(%r12)
  movq $0, (%r12)
  # <<<<< at (35,33)
  addq $40, %r12
  # [>>+<<-] at (35,38)
  movq (%r12), %rax
  addq %rax, -16(%r12)
  movq $0, (%r12)
  # + at (35,46)
  addq $1, (%r12)
  # < at (35,47)
  addq $8, %r12
  # [->-<] at (35,48)
  movq (%r12), %rax
  subq %rax, -8(%r12)
  movq $0, (%r12)
  # > at (35,54)
  subq $8, %r12
  # [ at (35,55)
  .loop32:
  cmpq $0, (%r12)
  je .break32
  # . at (35,58)
  movq -16(%r12), %rdi
  call putchar
  # << at (35,56)
  addq $16, %r12
  # [ at (35,63)
  .loop33:
  cmpq $0, (%r12)
  je .break33
  # + at (35,64)
  addq $1, (%r12)
  # . at (35,65)
  movq (%r12), %rdi
  call putchar
  # [-] at (35,66)
  movq $0, (%r12)
  # ] at (35,69)
  jmp .loop33
  .break33:
  # >> at (35,70)
  subq $16, %r12
  # - at (35,72)
  subq $1, (%r12)
  # ] at (35,73)
  jmp .loop32
  .break32:
  # > at (35,74)
  subq $8, %r12
  # [ at (35,75)
  .loop34:
  cmpq $0, (%r12)
  je .break34
  # . at (35,78)
  movq -16(%r12), %rdi
  call putchar
  # - at (36,1)
  subq $1, (%r12)
  # ] at (36,2)
  jmp .loop34
  .break34:
  # [-] at (36,4)
  movq $0, -8(%r12)
  # [-] at (36,8)
  movq $0, -16(%r12)
  # >>>>> at (36,3)
  subq $40, %r12
  # [ at (36,14)
  .loop35:
  cmpq $0, (%r12)
  je .break35
  # >> at (36,15)
  subq $16, %r12
  # [<<<<<<<<+>>>>>>>>-] at (36,17)
  movq (%r12), %rax
  addq %rax, 64(%r12)
  movq $0, (%r12)
  # << at (36,37)
  addq $16, %r12
  # - at (36,39)
  subq $1, (%r12)
  # ]] at (36,40)
  jmp .loop35
  .break35:
  # [-] at (36,44)
  movq $0, -16(%r12)
  # [-] at (36,50)
  movq $0, 8(%r12)
  # <<<<<<<<< at (36,42)
  addq $72, %r12
  # ] at (36,61)
  jmp .loop12
  .break12:
  # ++++++++++ at (36,62)
  addq $10, (%r12)
  # . at (36,72)
  movq (%r12), %rdi
  call putchar
  xorq %rax, %rax
  movq %rbp, %rsp
  popq %rbp
  ret
.section .note.GNU-stack
