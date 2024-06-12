package main

import (
	"compiler/pkg/lexer"
	"fmt"
)

func main() {
	input := `
	int main() {
	    int x = 10;
	    int y = 20;
	    int sum = x + y;
	    if (sum > 20) {
	        return 1;
	    } else {
	        return 0;
	    }
	}
	`

	l := lexer.New(input)

	for t := l.NextToken(); t.Type != lexer.EOF; t = l.NextToken() {
		fmt.Printf("Token Type: %s, Literal: %s\n", t.Type, t.Literal)
	}
}
