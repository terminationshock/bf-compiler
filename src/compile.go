package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func CompileAndLink(asm, c, file, mpiIncludePath, mpiLibPath string, hasMpi, outputAsm, verbose bool) error {
	fnameAsmo, err := compileAsm(asm, file, outputAsm, verbose)
	if fnameAsmo != "" {
		defer os.Remove(fnameAsmo)
	}
	if err != nil {
		return err
	}
	fnameCo, err := compileMpi(c, mpiIncludePath, hasMpi, verbose)
	if fnameCo != "" {
		defer os.Remove(fnameCo)
	}
	if err != nil {
		return err
	}

	return link(fnameAsmo, fnameCo, file, mpiLibPath, hasMpi, verbose)
}

func compileAsm(asm, file string, outputAsm, verbose bool) (string, error) {
	var fasm *os.File
	var err error

	if outputAsm {
		fasm, err = os.Create(file + ".asm")
	} else {
		fasm, err = os.CreateTemp("", "*.asm")
		defer os.Remove(fasm.Name())
	}
	if err != nil {
		return "", err
	}

	defer fasm.Close()

	_, err = fasm.Write([]byte(asm))
	if err != nil {
		return "", err
	}
	fasm.Close()

	fo, err := os.CreateTemp("", "*.o")
	if err != nil {
		return "", err
	}
	fo.Close()

	cmd := exec.Command("nasm", "-f", "elf64", "-o", fo.Name(), fasm.Name())
	if verbose {
		fmt.Println(cmd.String())
	}
	err = executeCommand(cmd)
	if err != nil {
		return fo.Name(), err
	}

	return fo.Name(), nil
}

func compileMpi(c, mpiIncludePath string, hasMpi, verbose bool) (string, error) {
	fc, err := os.CreateTemp("", "*.c")
	defer fc.Close()
	defer os.Remove(fc.Name())

	_, err = fc.Write([]byte(c))
	if err != nil {
		return "", err
	}
	fc.Close()

	fo, err := os.CreateTemp("", "*.o")
	if err != nil {
		return "", err
	}
	fo.Close()

	tokens := []string{"-c", "-o", fo.Name(), fc.Name()}
	if hasMpi {
		tokens = append(tokens, "-I", mpiIncludePath)
	}
	cmd := exec.Command("gcc", tokens...)
	if verbose {
		fmt.Println(cmd.String())
	}
	err = executeCommand(cmd)
	if err != nil {
		return fo.Name(), err
	}

	return fo.Name(), nil
}

func link(fnameAsmo, fnameCo, file, mpiLibPath string, hasMpi, verbose bool) error {
	tokens := []string{"-o", file, fnameAsmo, fnameCo}
	if hasMpi {
		tokens = append(tokens, "-L", mpiLibPath, "-lmpi", "-Wl,-rpath," + mpiLibPath)
	}
	cmd := exec.Command("gcc", tokens...)
	if verbose {
		fmt.Println(cmd.String())
	}
	err := executeCommand(cmd)
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
