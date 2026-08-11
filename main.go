package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "read error:", err)
				os.Exit(1)
			}
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", words[0])
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
