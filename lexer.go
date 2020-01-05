package lexer

import (
	"bufio"
	"fmt"
	"log"
)

type TokenType int

// Predefined Tokens that all callers must have
const (
	TokError TokenType = iota
	TokEof
)

type Token struct {
	Type TokenType
	Val  string
}

type StateFn func(*Lexer) StateFn

type Lexer struct {
	tokenCh    chan Token
	src        *bufio.Reader
	startState StateFn
	GotEof     bool
	tokNames   []string
}

const EofRune = -1

func (t Token) String() string {
	switch t.Type {
	case TokEof:
		return "EOF"
	case TokError:
		return t.Val
	}
	if len(t.Val) > 10 {
		return fmt.Sprintf("%.10q...", t.Val)
	}
	return fmt.Sprintf("<%d, %q>", t.Type, t.Val)
}

func NewLexer(ch chan Token, src *bufio.Reader, start StateFn, names []string) *Lexer {
	l := &Lexer{
		tokenCh:    ch,
		src:        src,
		startState: start,
		GotEof:     false,
		tokNames:   names,
	}
	return l
}

func (l *Lexer) Next() rune {
	r, _, err := l.src.ReadRune()
	if err != nil {
		if err.Error() == "EOF" {
			l.GotEof = true
			return EofRune
		}
		log.Panic("ReadRune got error")
	}
	return r
}

func (l *Lexer) UnNext() {
	err := l.src.UnreadRune()
	if err != nil {
		log.Panic("UnreadRune got error")
	}
}

func (l *Lexer) Peek() rune {
	r := l.Next()
	l.UnNext()
	return r
}

func (l *Lexer) Emit(t Token) {
	l.tokenCh <- t
}

func (l *Lexer) Run() {
	fmt.Println("lexer running...")
	for state := l.startState; state != nil; {
		state = state(l)
	}
	close(l.tokenCh)
}
