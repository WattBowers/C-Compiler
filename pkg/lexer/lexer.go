package lexer

type TokenType string

const (
	IDENTIFIER TokenType = "IDENTIFIER"
	KEYWORD    TokenType = "KEYWORD"
	NUMBER     TokenType = "NUMBER"
	OPERATOR   TokenType = "OPERATOR"
	DELIMITER  TokenType = "DELIMITER"
	EOF        TokenType = "EOF"
	ILLEGAL    TokenType = "ILLEGAL"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"int":      KEYWORD,
	"return":   KEYWORD,
	"if":       KEYWORD,
	"else":     KEYWORD,
	"auto":     KEYWORD,
	"double":   KEYWORD,
	"struct":   KEYWORD,
	"break":    KEYWORD,
	"long":     KEYWORD,
	"switch":   KEYWORD,
	"case":     KEYWORD,
	"enum":     KEYWORD,
	"register": KEYWORD,
	"typedef":  KEYWORD,
	"char":     KEYWORD,
	"extern":   KEYWORD,
	"union":    KEYWORD,
	"const":    KEYWORD,
	"float":    KEYWORD,
	"short":    KEYWORD,
	"unsigned": KEYWORD,
	"continue": KEYWORD,
	"for":      KEYWORD,
	"signed":   KEYWORD,
	"void":     KEYWORD,
	"default":  KEYWORD,
	"goto":     KEYWORD,
	"sizeof":   KEYWORD,
	"volatile": KEYWORD,
	"do":       KEYWORD,
	"static":   KEYWORD,
	"while":    KEYWORD,
}

type Lexer struct {
	input     string //whole string
	index     int    // index of position in string
	nextIndex int    // index of next position in string
	char      byte   // content of current index in string
}

// constructor
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if len(lexer.input) <= lexer.nextIndex {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.nextIndex]
	}
	lexer.nextIndex++
	lexer.index = lexer.nextIndex - 1
}

func (lexer *Lexer) NextToken() Token {

	var t Token
	lexer.skipWhiteSpace()

	switch lexer.char {
	case '=':
		t = Token{Type: OPERATOR, Literal: "="}
	case '+':
		t = Token{Type: OPERATOR, Literal: "+"}
	case '-':
		t = Token{Type: OPERATOR, Literal: "-"}
	case '*':
		t = Token{Type: OPERATOR, Literal: "*"}
	case '/':
		t = Token{Type: OPERATOR, Literal: "/"}
	case '>':
		t = Token{Type: OPERATOR, Literal: ">"}
	case '<':
		t = Token{Type: DELIMITER, Literal: "<"}
	case ';':
		t = Token{Type: DELIMITER, Literal: ";"}
	case '(':
		t = Token{Type: DELIMITER, Literal: "("}
	case ')':
		t = Token{Type: DELIMITER, Literal: ")"}
	case '{':
		t = Token{Type: DELIMITER, Literal: "{"}
	case '}':
		t = Token{Type: DELIMITER, Literal: "}"}
	case 0:
		t = Token{Type: EOF, Literal: ""}

	default:
		if isLetter(lexer.char) {
			t.Literal = lexer.findKeyword()
			t.Type = lexer.lookupIdentity(t.Literal)
			return t
		} else if isNumber(lexer.char) {
			return Token{Type: NUMBER, Literal: lexer.findNumber()}
		} else {
			return Token{Type: ILLEGAL, Literal: string(lexer.char)}
		}
	}
	lexer.readChar()
	return t

}

func isLetter(char byte) bool {
	if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' {
		return true
	} else {
		return false
	}
}

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	} else {
		return false
	}
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.char == ' ' || lexer.char == '\n' || lexer.char == '\t' || lexer.char == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) lookupIdentity(literal string) TokenType {
	if _, exists := keywords[literal]; exists {
		return KEYWORD
	} else {
		return IDENTIFIER
	}
}

func (lexer *Lexer) findKeyword() string {
	index := lexer.index
	for isLetter(lexer.char) {
		lexer.readChar()
	}
	return lexer.input[index:lexer.index]
}

func (lexer *Lexer) findNumber() string {
	index := lexer.index
	for isNumber(lexer.char) {
		lexer.readChar()
	}
	return lexer.input[index:lexer.index]
}
