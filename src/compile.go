package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Compile(asm string, file string, mpiPath string, outputAsm bool, verbose bool) error {
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


	fo, err := os.CreateTemp("", "*.o")
    if err != nil {
		return err
    }
    fo.Close()
	defer os.Remove(fo.Name())

	cmd := exec.Command("nasm", "-f", "elf64", "-o", fo.Name(), fasm.Name())
	if verbose {
		fmt.Println(cmd.String())
	}
	err = executeCommand(cmd)
	if err != nil {
		return err
	}

	cmd = exec.Command("gcc", "-o", file, "-L", mpiPath, "-lmpi", "-Wl,-rpath," + mpiPath, fo.Name())
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
