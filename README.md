# Atcoder Testcase Runner

## Description
A simple cli tool to run testcases for Atcoder problems.
Fetches and saves testcases from www.atcoder.jp and runs them against the solution file.
Currently, supports:
- C++
- Java
- Kotlin
- Python

## Installation
```
git clone https://www.github.com/arkzuse/atcoder-testcase-runner.git
```
```
cd atcoder-testcase-runner
```
Install dependencies
```
go mod tidy
```
Build the binary
```
go build
```

## Usage
- You can add commands for your language in `codeRunner.go`. In future, a config file will be added for custom command and testcase files.
- Code should be in atcoder.* file.
- If `-c` or `-t` flags are provided then testcases are fetched from task page otherwise previous testcases are used.

To run from source:
```
go run main.go atcoder.* -c <contest> -t <task>
```

For help:
```
./atcoder-testcase-runner -h
```

## Example
```
./atcoder-testcase-runner atcoder.py -c abd371 -t b
```
