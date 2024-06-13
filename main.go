package main

import (
	"fmt"
	"os"

	inputs "ascii-art/pkg"
)

func main() {
	input, _ := inputs.InputArgs(os.Args)
	asciiFields := inputs.FileChoice(os.Args)
	printWords(input, asciiFields)
}

// printing of the characters
func printWords(input []string, asciiFields []string) {
	for _, word := range input {
		if word == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if !validChar(char) {
						return
					}
					startPoint := Start(int(char))
					fmt.Print(asciiFields[startPoint+i])
				}
				fmt.Println()
			}
		}
	}
}

// starting positions of the characters
func Start(s int) int {
	pos := int(s-32)*9 + 1
	return pos
}

func validChar(s rune) bool {
	if !(s >= ' ' && s <= '~') {
		fmt.Println(string(s) + " " + "is not valid character")
		return false
	}
	return true
}
