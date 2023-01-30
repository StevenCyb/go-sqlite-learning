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

		switch input {
		case ".exit":
			return
		default:
			fmt.Printf("Unrecognized command '%s'.\n", input)
		}
	}
}
