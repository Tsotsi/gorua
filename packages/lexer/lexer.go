package lexer

import (
	"github.com/pkg/errors"
	"io"
	"unicode"
	"xlin/learn-llvm/packages/token"
)

type Lexer struct {
	input     io.RuneScanner
	line      int
	cur       token.Pos
	TokenPipe chan token.Token
}

func New(ipt io.RuneScanner) *Lexer {
	return &Lexer{
		input:     ipt,
		TokenPipe: make(chan token.Token),
		cur:       token.NoPos,
	}
}

func (l Lexer) NextToken() token.Token {
	var r rune
	var s int
	var e error
	for {
		r, s, e = l.input.ReadRune()
		l.cur += token.Pos(s)
		if e != nil {
			if errors.Is(e, io.EOF) {
				return token.EOF
			}
			return token.ILLEGAL
		}
		if unicode.IsSpace(r) {
			continue
		}
	}
	return token.ILLEGAL
}
