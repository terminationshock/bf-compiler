package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func CompileAndLink(assembly, exeFile string, keepAssemblyFile, verbose bool) error {
	var fs *os.File
	var err error

	if keepAssemblyFile {
		fs, err = os.Create(exeFile + ".s")
	} else {
		fs, err = os.CreateTemp("", "*.s")
		defer os.Remove(fs.Name())
	}
	if err != nil {
		return err
	}

	defer fs.Close()

	_, err = fs.Write([]byte(assembly))
	if err != nil {
		return err
	}
	if verbose {
		fmt.Println("Assembly written to", fs.Name())
	}
	fs.Close()

	cmd := exec.Command("gcc", "-o", exeFile, fs.Name())
	if verbose {
		fmt.Println(cmd.String())
	}
	return executeCommand(cmd, verbose)
}

func executeCommand(cmd *exec.Cmd, verbose bool) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	outputStdout, _ := io.ReadAll(stdout)
	outputStderr, _ := io.ReadAll(stderr)
	err = cmd.Wait()

	if verbose {
		os.Stdout.Write(outputStdout)
	}
	if verbose || err != nil {
		os.Stderr.Write(outputStderr)
	}
	return err
}
