# Brainfuck Compiler

This repository hosts a [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) compiler for Linux x86-64 systems written in Go. It translates the source file to GNU highly-optimized assembly code and compiles it using GCC.

## Usage

Run `make` to compile the compiler itself. Alternatively, download the binary of the latest release from [here](https://github.com/terminationshock/bf-compiler/releases/latest).

You can then compile a Brainfuck source file by invoking `bf <FILE>`. Several command-line options exist, which can be viewed with `bf --help`.

## Run tests

With `make check`, a suite of test programs will be compiled and run. Most of the tests are taken from [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck), while [one test](test/pi.bf) has been created by [Felix Nawothnig](mailto:felix.nawothnig@t-online.de).
