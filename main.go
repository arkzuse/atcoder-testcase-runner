package main

import (
	"atcoder-testcase-runner/utils"
	"fmt"
)

func main() {
	var samples = utils.ScrapeTestcase("abc371", "b")

	for _, sample := range samples {
		fmt.Println(sample[0])
		fmt.Println(sample[1])
		fmt.Println()
	}
}
