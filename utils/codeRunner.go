package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
)

func RunSolution(fileName string, input string) (string, error) {
	os := runtime.GOOS
	command, err := generateCommand(fileName)
	if err != nil {
		return "", err
	}

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
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s: %s", err, stderr.String())
	}

	return out.String(), nil
}

func generateCommand(fileName string) (string, error) {
	ext := filepath.Ext(fileName)

	switch ext {
	case ".cpp":
		return fmt.Sprintf("g++ -lm %s -o %s && ./%s", fileName, fileName[:len(fileName)-len(ext)], fileName[:len(fileName)-len(ext)]), nil
	case ".java":
		return fmt.Sprintf("javac %s && java %s", fileName, fileName[:len(fileName)-len(ext)]), nil
	case ".kt":
		return fmt.Sprintf("kotlinc %s -include-runtime -d %s.jar && java -jar %s.jar", fileName, fileName[:len(fileName)-len(ext)], fileName[:len(fileName)-len(ext)]), nil
	case ".py":
		return fmt.Sprintf("python3 %s", fileName), nil
	default:
		return "", fmt.Errorf("unsupported file extension")
	}
}
