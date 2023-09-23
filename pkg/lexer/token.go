package lexer

import "fmt"

type Token struct {
	Type    Type
	Literal Literal
}

type (
	Type    = string
	Literal = []rune
)

const (
	eof      Type = "eof"
	invalid       = "invalid"
	comma         = "comma"
	colon         = "colon"
	lBrace        = "lBrace"
	rBrace        = "rBrace"
	lBracket      = "lBracket"
	rBracket      = "rBracket"
	str           = "string"
	integer       = "integer"
)

func NewToken(t Type, literal string) Token {
	return Token{Type: t, Literal: []rune(literal)}
}

func (t *Token) String() string {
  return fmt.Sprintf("token::%s %s", string(t.Type), string(t.Literal))
}
