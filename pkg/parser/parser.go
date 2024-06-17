package parser

import (
	"compiler/pkg/lexer"
)

type node interface{}

type numberNode struct {
	Num int
}

type operatorNode struct {
	Left     node
	Operator string
	Right    node
}

type parser struct {
	lexer  *lexer.Lexer
	tokens []lexer.Token
}

// initialize parser
func New(lexer *lexer.Lexer) *parser {
	p := &parser{lexer: lexer}
	p.token = p.lexer.NextToken()
	return p
}

func (p *parser) Parse() (node, error) {

}
