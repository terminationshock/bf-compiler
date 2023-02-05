package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Compile(asm, file string) error {
	fasm, err := os.CreateTemp("", "*.asm")
    if err != nil {
		return err
    }

    defer fasm.Close()
	defer os.Remove(fasm.Name())

	fo, err := os.CreateTemp("", "*.o")
    if err != nil {
		return err
    }

    defer fo.Close()
	defer os.Remove(fo.Name())

	_, err = fasm.Write([]byte(asm))
	if err != nil {
		return err
	}
	fasm.Close()

	cmd := exec.Command("nasm", "-f", "elf64", "-o", fo.Name(), fasm.Name())
	err = executeCommand(cmd)
	if err != nil {
		return err
	}

	cmd = exec.Command("gcc", "-o", file, fo.Name())
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
		fmt.Printf("%s\n", output)
		return err
	}

	return nil
}
