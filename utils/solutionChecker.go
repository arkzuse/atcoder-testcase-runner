package utils

func CheckDiff(output string, solution string) bool {
	if output == solution {
		return true
	}

	return false
}
