package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hculpan/klang/lexer"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: lexer <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	input := strings.TrimSpace(string(content))
	tokens, err := lexer.Tokenize(input)
	if err != nil {
		fmt.Printf("Error tokenizing: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Tokens:")
	for _, token := range tokens {
		fmt.Println(token)
	}
}
