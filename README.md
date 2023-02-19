# Brainfuck Compiler

This repository hosts a [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) compiler for Linux x86-64 systems written in Go. It translates the source file to highly-optimized GNU assembly code and compiles it using GCC.

## Usage

Run `make` to compile the compiler itself. Alternatively, download the binary of the latest release from [here](https://github.com/terminationshock/bf-compiler/releases/latest).

You can then compile a Brainfuck source file by invoking `bf <FILE>`. Several command-line options exist, which can be viewed with `bf --help`.

## Run tests

With `make check`, a suite of test programs will be compiled and run. Most of the tests are taken from [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck), while [one test](test/pi.bf) has been created by Felix Nawothnig.

## Benchmarks

The following test case has been used for benchmarking the compiler: `factor.b` by [Brian Raiter](https://github.com/BR903/ELFkickers/blob/master/ebfc/bf) with input `6543210987654321`. The code has been compiled with different compilers and optimization levels.

| Compiler | Runtime |
| -------- | ------- |
| [ebfc (commit e7fba94)](https://github.com/BR903/ELFkickers/tree/e7fba942df51e756897224cff5aa853de8fafd90/ebfc) | 23 seconds |
| [bfc (verion 1.7.0)](https://bfc.wilfred.me.uk/) | 8.5 seconds |
| this repository | 7.8 seconds |
