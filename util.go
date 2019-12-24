package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// IsExist check if the directory exist
// @path string Full directory path
// @return bool, error
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

// IsEmpty check if the directory is empty
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}

	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

// ExecCmd execute a shell command
func ExecCmd(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	fmt.Printf("Running command: %s\n%s\n%s\n", command, stdout.String(), stderr.String())
	return stdout.String(), stderr.String(), err
}

// ExecSlowCmd can run a slow command and capture its output
func ExecSlowCmd(command string) {
	fmt.Printf("Running command: %s\n", command)
	cmd := exec.Command("bash", "-c", command)

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("%s\n%s\n", outStr, errStr)
}

func Copy(src, dest string) error {
	sourceContent, _ := Asset(src)

	if lastIndex := strings.LastIndex(dest, "/"); lastIndex != -1 {
		destPath := dest[:lastIndex]
		if isExist, _ := IsExist(destPath); isExist == false {
			os.MkdirAll(destPath, os.ModePerm)
		}
	}

	targetFile := project.Directory + string(os.PathSeparator) + Path(dest)
	destination, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = destination.Write(sourceContent)
	return err
}

func Path(path string) string {
	return strings.Replace(path, "/", string(os.PathSeparator), -1)
}
