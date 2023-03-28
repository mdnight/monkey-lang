package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tokType token.TokenType
	var literal string

	l.skipWhitespace()

	switch l.ch {
	// Operators
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tokType = token.EQ
			literal = "=="
		} else {
			tokType = token.ASSIGN
			literal = "="
		}
	case '+':
		tokType = token.PLUS
		literal = "+"
	case '-':
		tokType = token.MINUS
		literal = "-"
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tokType = token.NEQ
			literal = "!="
		} else {
			tokType = token.BANG
			literal = "!"
		}
	case '*':
		tokType = token.ASTERISK
		literal = "*"
	case '/':
		tokType = token.SLASH
		literal = "/"
	case '<':
		tokType = token.LT
		literal = "<"
	case '>':
		tokType = token.GT
		literal = ">"
	case ';':
		tokType = token.SEMICOLON
		literal = ";"
	case ',':
		tokType = token.COMMA
		literal = ","
	case '(':
		tokType = token.LPAREN
		literal = "("
	case ')':
		tokType = token.RPAREN
		literal = ")"
	case '{':
		tokType = token.LBRACE
		literal = "{"
	case '}':
		tokType = token.RBRACE
		literal = "}"
	case 0x00:
		tokType = token.EOF
		literal = ""
	default:
		if isLetter(l.ch) {
			literal = l.readIdentifier()
			tokType = token.LookupIndent(literal)
		} else if isDigit(l.ch) {
			literal = l.readNumber()
			tokType = token.INT
		} else {
			literal = string(l.ch)
			tokType = token.ILLEGAL
		}
		return newToken(tokType, literal)
	}

	l.readChar()
	return newToken(tokType, literal)
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() (literal string) {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() (literal string) {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
	}
}
