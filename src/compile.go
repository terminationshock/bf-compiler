package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Compile(asm, c, file, mpiIncludePath, mpiLibPath string, outputAsm bool, verbose bool) error {
	var fasm *os.File
	var err error

	if outputAsm {
		fasm, err = os.Create(file + ".asm")
	} else {
		fasm, err = os.CreateTemp("", "*.asm")
		defer os.Remove(fasm.Name())
	}
	if err != nil {
		return err
	}

	defer fasm.Close()

	_, err = fasm.Write([]byte(asm))
	if err != nil {
		return err
	}
	fasm.Close()

	fc, err := os.CreateTemp("", "*.c")
	defer fc.Close()
	defer os.Remove(fc.Name())

	_, err = fc.Write([]byte(c))
	if err != nil {
		return err
	}
	fc.Close()

	foAsm, err := os.CreateTemp("", "*.o")
	if err != nil {
		return err
	}
	foAsm.Close()
	defer os.Remove(foAsm.Name())

	foC, err := os.CreateTemp("", "*.o")
	if err != nil {
		return err
	}
	foC.Close()
	defer os.Remove(foC.Name())

	cmd := exec.Command("nasm", "-f", "elf64", "-o", foAsm.Name(), fasm.Name())
	if verbose {
		fmt.Println(cmd.String())
	}
	err = executeCommand(cmd)
	if err != nil {
		return err
	}

	cmd = exec.Command("gcc", "-c", "-o", foC.Name(), "-I", mpiIncludePath, fc.Name())
	if verbose {
		fmt.Println(cmd.String())
	}
	err = executeCommand(cmd)
	if err != nil {
		return err
	}

	cmd = exec.Command("gcc", "-o", file, "-L", mpiLibPath, "-lmpi", "-Wl,-rpath," + mpiLibPath, foAsm.Name(), foC.Name())
	if verbose {
		fmt.Println(cmd.String())
	}
	err = executeCommand(cmd)
	if err != nil {
		return err
	}

	return nil
}

func executeCommand(cmd *exec.Cmd) error {
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	output, _ := io.ReadAll(stderr)
	err = cmd.Wait()
	if err != nil {
		os.Stderr.Write(output)
		return err
	}

	return nil
}
