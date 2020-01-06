package lexer

import (
	"bufio"
	"fmt"
	"log"
	"strings"
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

type charTestFn func(rune) bool

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

func (l *Lexer) AcceptRun(testFunc charTestFn) string {
    var b strings.Builder
    for {
        r := l.Next()
        if testFunc(r) {
            b.WriteRune(r)
        } else {
            // We have hit the end of the string of chars we want
            // If we just got EOF, don't try to push it back
            if r != EofRune {
                l.UnNext()
            }
            break
        }
    }
    return b.String()
}

