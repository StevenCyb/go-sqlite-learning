package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("db> ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")

		// meta-commands
		if input[0] == '.' {
			switch input {
			case ".exit":
				return
			default:
				fmt.Printf("Unrecognized command '%s'.\n", input)
			}

			continue
		}

		// statements
		switch {
		case strings.HasPrefix(input, "insert "):
			fmt.Println("This is where we would do an insert.")
		case strings.HasPrefix(input, "select "):
			fmt.Println("This is where we would do an select.")
		default:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", input)
		}
	}
}
