package lexer

import (
	"bufio"
	"fmt"
	"log"
)

type TokenType int

const (
	TokError TokenType = iota
	TokEof
)

type Token struct {
	// typ TokenType
	// val string
	Type TokenType
	Val  string
}

type StateFn func(*Lexer) StateFn

type Lexer struct {
	tokenCh    chan Token
	src        *bufio.Reader
	startState StateFn
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

func NewLexer(ch chan Token, src *bufio.Reader, start StateFn) *Lexer {
	l := &Lexer{
		tokenCh:    ch,
		src:        src,
		startState: start,
	}
	return l
}

func (l *Lexer) Next() rune {
	r, _, err := l.src.ReadRune()
	if err != nil {
		if err.Error() == "EOF" {
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

func (l *Lexer) Emit(t Token) {
	l.tokenCh <- t
}

func (l *Lexer) Run() {
	fmt.Println("lexer running...")
	for state := l.startState; state != nil; {
		state = state(l)
	}
	/***
		for {
			r, _, err := l.src.ReadRune()
			if  err != nil {
				// This happens at EOF
				fmt.Println("error", err)
				if err.Error() == "EOF" {
					fmt.Println("got EOF")
				}
				break
			}
			fmt.Printf("lexer read: %v %q\n", r, r)
			l.tokenCh <- Token{r: r, val: string(r)}
		}
	***/
	close(l.tokenCh)
}
