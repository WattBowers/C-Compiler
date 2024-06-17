package main

import (
	"compiler/pkg/lexer"
	"fmt"
	"os"
)

func main() {

	bytes, _ := os.ReadFile("./examples/02.lang")

	token := lexer.New(string(bytes))

	var tokenList []lexer.Token

	for t := token.NextToken(); t.Type != lexer.EOF; t = token.NextToken() {
		tokenList = append(tokenList, t)
	}

	for _, token := range tokenList {
		fmt.Printf("Type: %-10s Literal: %-10s\n", token.Type, token.Literal)
	}
}
