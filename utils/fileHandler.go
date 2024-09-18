package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WriteInputFile(fileName string, contest string, task string, input []string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	content := inputFileContent(contest, task, input)
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func ReadInputFile(fileName string) ([]string, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	var input []string

	defer file.Close()

	// read whole file
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return nil, err
	}

	// convert bytes to string
	str := string(bs)

	var elem string
	lines := strings.Split(str, "\n")
	inputSize, err := strconv.Atoi(strings.TrimSpace(lines[2]))
	var count = 0
	if err != nil {
		return nil, fmt.Errorf("failed to parse input sample size")
	}
	for _, l := range lines {
		if count == inputSize+2 {
			break
		}
		if l == "" {
			count++
			input = append(input, elem)
			elem = ""
		} else {
			elem += l + "\n"
		}
	}
	if elem != "" {
		input = append(input, elem)
	}

	return input, nil
}

func inputFileContent(contest string, task string, input []string) string {
	var content string

	content += fmt.Sprintf("%s %s\n\n", contest, task)
	content += fmt.Sprintf("%d\n\n", len(input))
	for _, i := range input {
		content += fmt.Sprintf("%s\n", i)
	}

	return content
}
