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

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tokType = token.ASSIGN
	case ';':
		tokType = token.SEMICOLON
	case ',':
		tokType = token.COMMA
	case '+':
		tokType = token.PLUS
	case '(':
		tokType = token.LPAREN
	case ')':
		tokType = token.RPAREN
	case '{':
		tokType = token.LBRACE
	case '}':
		tokType = token.RBRACE
	case 0x00:
		tokType = token.EOF
	default:
		var tok token.Token
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokType = token.LookupIndent(literal)
			tok = newToken(tokType, literal)
		} else if isDigit(l.ch) {
			literal := l.readNumber()
			tok = newToken(token.INT, literal)
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch))
		}

		return tok
	}

	var ch string
	if tokType == token.EOF {
		ch = ""
	} else {
		ch = string(l.ch)
	}

	l.readChar()
	return newToken(tokType, ch)
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
