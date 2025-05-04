package utils

import (
	"strings"
)

func CleanInput(input string) []string {
	trimmedInput := strings.Trim(input, " ")
	args := strings.SplitAfter(trimmedInput, " ")[1:]
	return args
}
