package utils

import (
	"fmt"
	"os"
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

func inputFileContent(contest string, task string, input []string) string {
	var content string

	content += fmt.Sprintf("%s %s\n\n", contest, task)
	content += fmt.Sprintf("%d\n\n", len(input))
	for _, i := range input {
		content += fmt.Sprintf("%s\n", i)
	}

	return content
}
