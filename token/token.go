package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	// Operators
	ASSIGN   TokenType = "ASSIGN"
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	BANG     TokenType = "BANG"
	ASTERISK TokenType = "ASTERISK"
	SLASH    TokenType = "SLASH"

	EQ  TokenType = "EQ"
	NEQ TokenType = "NEQ"
	LT  TokenType = "LT"
	GT  TokenType = "GT"

	// Delimiters
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	LBRACE TokenType = "LBRACE"
	RBRACE TokenType = "RBRACE"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	RETURN   TokenType = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIndent(literal string) TokenType {
	if tok, ok := keywords[literal]; ok {
		return tok
	}

	return IDENT
}
