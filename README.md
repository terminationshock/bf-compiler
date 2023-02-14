# Brainfuck Compiler

This repository hosts a [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) compiler for Linux x86-64 systems written in Go. It translates the source file to GNU highly-optimized assembly code and compiles it using GCC.

## Usage

Run `make` to compile the compiler itself. You can then compile a Brainfuck source file by invoking `./bf <FILE>`. Several command-line options exist, which can be viewed with `./bf --help`.

## Run tests

With `make check`, a suite of test programs will be compiled and run.

## License

[GPL v3](https://www.gnu.org/licenses/gpl-3.0)
