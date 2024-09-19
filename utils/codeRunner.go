package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func RunSolution(command string, input string) (string, error) {
	os := runtime.GOOS
	var cmd *exec.Cmd
	switch os {
	case "darwin":
		cmd = exec.Command("sh", "-c", command)
	case "linux":
		cmd = exec.Command("sh", "-c", command)
	case "windows":
		cmd = exec.Command("cmd", "/c", command)
	default:
		return "", fmt.Errorf("unsupported OS")
	}

	cmd.Stdin = bytes.NewBufferString(input)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
