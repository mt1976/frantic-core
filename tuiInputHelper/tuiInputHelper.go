package tuiInputHelper

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func GetUserInput(question string, inputWidth int, inputMandatory bool, isPassword bool) string {
	var response string

	requestInput := func() string {
		fmt.Print("\033[1m" + question + "\033[0m ")

		if isPassword {
			// Enable raw mode and mask with '*'
			oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
			if err != nil {
				panic(err)
			}
			defer term.Restore(int(os.Stdin.Fd()), oldState)

			input := make([]rune, 0, inputWidth)

			for buf := make([]byte, 1); ; {
				os.Stdin.Read(buf)

				if char := buf[0]; char == '\r' || char == '\n' {
					// Enter key
					break
				} else if char == 127 || char == 8 {
					// Backspace
					if len(input) > 0 {
						input = input[:len(input)-1]
						fmt.Print("\b \b")
					}
				} else if len(input) < inputWidth && char >= 32 && char <= 126 {
					// Printable character
					input = append(input, rune(char))
					fmt.Print("*")
				}
			}
			fmt.Println()
			return string(input)
		} else {
			// Normal input
			var input string
			fmt.Scanln(&input)
			if len(input) > inputWidth {
				input = input[:inputWidth]
			}
			return strings.TrimSpace(input)
		}
	}

	// First input attempt
	response = requestInput()

	// Repeat if mandatory and empty
	if inputMandatory {
		for response == "" {
			fmt.Printf("\n\033[1m%*s %s is \033[5mMANDATORY\033[0m\n", inputWidth+1, " ", question)
			response = requestInput()
		}
	}

	fmt.Println()
	return response
}
