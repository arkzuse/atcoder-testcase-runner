package main

import (
	"atcoder-testcase-runner/utils"
	"fmt"
)

func main() {
	samples, err := utils.ScrapeTestcase("abc371", "a")

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sample := range samples {
		fmt.Println(sample[0])
		fmt.Println(sample[1])
		fmt.Println()
	}
}
