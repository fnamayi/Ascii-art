package inputs

import (
	"fmt"
	"os"
	"strings"
)

var args string

// InputArgs function processes command-line arguments
func InputArgs(osArgs []string) ([]string, int) {
	if !(len(osArgs) == 2 || len(osArgs) == 3) {
		fmt.Println("usage: go run . <input string> <optional flag>")
		os.Exit(0)
	} else if len(osArgs) == 3 && !(osArgs[2] == "-t" || osArgs[2] == "-sh") {
		fmt.Println("Only one string is accepted")
		os.Exit(0)
	} else {
		args = osArgs[1]
	}

	args = strings.ReplaceAll(args, "\n", "\\n")
	args = strings.ReplaceAll(args, "\\t", " ")

	input := strings.Split(args, "\\n")

	if input[0] == "" && len(input) == 1 {
		os.Exit(0)
	} else if input[0] == "" && input[1] == "" && len(input) == 3 {
		fmt.Println()
		os.Exit(0)
	}
	return input, len(os.Args)
}

// FileChoice function selects the appropriate ASCII art file to read
func FileChoice(osArgs []string) []string {
	var file []byte
	var asciiFields []string
	// split the file
	if len(osArgs) == 3 && osArgs[2] == "-t" {
		filec, err := os.ReadFile("thinkertoy.txt")
		file = filec
		asciiField := strings.Split(string(file), "\r\n")
		asciiFields = asciiField
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if len(osArgs) == 3 && osArgs[2] == "-sh" {
		filec, err := os.ReadFile("shadow.txt")
		file = filec
		asciiField := strings.Split(string(file), "\n")
		asciiFields = asciiField
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		filec, err := os.ReadFile("standard.txt")
		file = filec
		asciiField := strings.Split(string(file), "\n")
		asciiFields = asciiField
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if len(string(file)) == 0 {
		fmt.Println("Empty file")
		os.Exit(0)
	}

	if len(asciiFields) != 856 {
		fmt.Println("The File Banner used has been tampered with")

		os.Exit(0)
	}
	return asciiFields
}
