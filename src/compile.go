package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func CompileAndLink(s, c, exeFile, mpiIncludePath, mpiLibPath string, hasMpi, keepAssemblyFile, verbose bool) error {
	fnameSo, err := compileAssembly(s, exeFile, keepAssemblyFile, verbose)
	if fnameSo != "" {
		defer os.Remove(fnameSo)
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

	return link(fnameSo, fnameCo, exeFile, mpiLibPath, hasMpi, verbose)
}

func compileAssembly(s, exeFile string, keepAssemblyFile, verbose bool) (string, error) {
	var fs *os.File
	var err error

	if keepAssemblyFile {
		fs, err = os.Create(exeFile + ".s")
	} else {
		fs, err = os.CreateTemp("", "*.s")
		defer os.Remove(fs.Name())
	}
	if err != nil {
		return "", err
	}

	defer fs.Close()

	_, err = fs.Write([]byte(s))
	if err != nil {
		return "", err
	}
	fs.Close()

	fo, err := os.CreateTemp("", "*.o")
	if err != nil {
		return "", err
	}
	fo.Close()

	cmd := exec.Command("gcc", "-c", "-o", fo.Name(), fs.Name())
	if verbose {
		fmt.Println(cmd.String())
	}
	err = executeCommand(cmd, verbose)
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
	err = executeCommand(cmd, verbose)
	if err != nil {
		return fo.Name(), err
	}

	return fo.Name(), nil
}

func link(fnameSo, fnameCo, exeFile, mpiLibPath string, hasMpi, verbose bool) error {
	tokens := []string{"-o", exeFile, fnameSo, fnameCo}
	if hasMpi {
		tokens = append(tokens, "-L", mpiLibPath, "-lmpi", "-Wl,-rpath," + mpiLibPath)
	}
	cmd := exec.Command("gcc", tokens...)
	if verbose {
		fmt.Println(cmd.String())
	}
	err := executeCommand(cmd, verbose)
	if err != nil {
		return err
	}

	return nil
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
	if err != nil {
		return err
	}

	return nil
}
