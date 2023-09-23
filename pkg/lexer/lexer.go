package lexer

import (
	"log"
	"unicode"
)

type Index = int

type Lexer struct {
	input []rune
}

func NewLexer(input string) *Lexer {
	l := Lexer{input: []rune(input)}
	return &l
}

func (l *Lexer) Tokenize() (tokens []Token) {
	for {
		tokens = append(tokens, l.ReadToken())
		if aboba := l.Next(); !aboba {
			break
		}
	}

	return
}

func (l *Lexer) ReadToken() Token {
	var t Token

	l.ignoreWhitespace()

	switch l.char() {
	case rune(0):
		t = NewToken(eof, string(l.char()))
	case ',':
		t = NewToken(comma, string(l.char()))
	case ':':
		t = NewToken(colon, string(l.char()))
	case '{':
		t = NewToken(lBrace, string(l.char()))
	case '}':
		t = NewToken(rBrace, string(l.char()))
	case '[':
		t = NewToken(lBracket, string(l.char()))
	case ']':
		t = NewToken(rBracket, string(l.char()))
	default:
		if ok, lit := l.isString(); ok {
			t = NewToken(str, string(lit))
		} else if i, lit := l.isInteger(); i {
			t = NewToken(integer, string(lit))
		} else {
			t = NewToken(invalid, string(l.char()))
		}
	}

	return t
}

func (l *Lexer) Next() bool {
	if l.char() != rune(0) {
		l.input = l.input[1:]
		return true
	}
	return false
}

func (l *Lexer) char() rune {
	if len(l.input) == 0 {
		return rune(0)
	}

	return l.input[0]
}

func (l *Lexer) isString() (check bool, literal []rune) {
	if l.char() != '"' {
		return false, nil
	}

	literal = append(literal, l.char())
	l.Next()

	for ; l.char() != rune(0) && l.char() != '"'; l.Next() {
		literal = append(literal, l.char())
	}

	if l.char() != '"' {
		log.Fatalf("unterminated string %q", string(literal))
	}

	literal = append(literal, l.char())
	l.Next()

	return true, literal
}

func (l *Lexer) isInteger() (ok bool, literal []rune) {
	if !unicode.IsDigit(l.char()) && string(l.char()) != "-" {
		return false, nil
	}

	for ; unicode.IsDigit(l.char()); l.Next() {
		literal = append(literal, l.char())
	}
	return true, literal
}

func (l *Lexer) ignoreWhitespace() {
	for {
		switch l.char() {
		case ' ', '\n', '\t', '\r':
			l.Next()
		default:
			return
		}
	}
}
