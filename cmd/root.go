package cmd

import (
	"atcoder-testcase-runner/utils"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var FileName string
var Contest string
var Task string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "atcoder-testcase-runner [file]",
	Short: "A tool to run all test cases of an AtCoder problem",
	Long: `A tool to run all test cases of an AtCoder problem.
Fetches input and answer from www.atcoder.jp and saves them in text files.
To run for a new problem or refresh inputs/solution just provide contest and tast flags.
Currently supports C++, Java, Kotlin and Python.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("File required")
		} else {
			FileName = args[0]
			fileInput, err := utils.ReadFile("atcoderInput.txt")
			if err != nil {
				cmd.Println(err)
			}
			fileAnswer, err := utils.ReadFile("atcoderAnswer.txt")
			if err != nil {
				cmd.Println(err)
			}

			// contest and task from file
			if Contest == "" && Task == "" {
				if fileInput[0] != fileAnswer[0] {
					cmd.Println("Contest name mismatch in input and answer file.\nFetch them again.")
					return
				}

				Contest = strings.Split(fileInput[0], " ")[0]
				Task = strings.TrimSpace(strings.Split(fileInput[0], " ")[1])
			} else {
				// check flags provided else use from file
				if Contest == "" {
					Contest = strings.Split(fileInput[0], " ")[0]
				}
				if Task == "" {
					Task = strings.TrimSpace(strings.Split(fileInput[0], " ")[1])
				}

				cmd.Println("Fetching test cases...")

				// Empty the fileInput and fileAnswer
				fileInput = []string{}
				fileAnswer = []string{}

				res, err := utils.ScrapeTestcase(Contest, Task)
				if err != nil {
					cmd.Printf("Failed to fetch test cases: %v\n", err)
					return
				}

				for i := 0; i < len(res); i++ {
					fileInput = append(fileInput, res[i][0])
					fileAnswer = append(fileAnswer, res[i][1])
				}

				err = utils.WriteFile("atcoderInput.txt", Contest, Task, fileInput)
				if err != nil {
					cmd.Printf("Failed to write input file: %v\n", err)
					return
				}
				err = utils.WriteFile("atcoderAnswer.txt", Contest, Task, fileAnswer)
				if err != nil {
					cmd.Printf("Failed to write answer file: %v\n", err)
					return
				}

				// prepend contest task and size
				fileInput = append([]string{Contest + " " + Task, strconv.Itoa(len(fileInput))}, fileInput...)
				fileAnswer = append([]string{Contest + " " + Task, strconv.Itoa(len(fileAnswer))}, fileAnswer...)
			}

			cmd.Println("Running test cases...")
			passed := 0
			var outputArray []string

			for i := 0; i < len(fileInput)-2; i++ {
				output, err := utils.RunSolution(FileName, fileInput[i+2])
				if utils.CheckDiff(output, fileAnswer[i+2]) {
					cmd.Println("Test case", i+1, "passed")
					passed++
				} else {
					cmd.Println("Test case", i+1, "failed")
				}

				if err != nil {
					cmd.Println(err)
				}

				outputArray = append(outputArray, output)

			}

			cmd.Printf("[%d/%s] test cases passed\n", passed, strings.TrimSpace(fileInput[1]))
			err = utils.WriteFile("atcoderOutput.txt", Contest, Task, outputArray)
			if err != nil {
				cmd.Println(err)
			}
			cmd.Println("Output written to atcoderOutput.txt")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&Contest, "contest", "c", "", "Contest name")
	rootCmd.Flags().StringVarP(&Task, "task", "t", "", "Task name")
}
