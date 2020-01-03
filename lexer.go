package lexer

import (
	"bufio"
	"fmt"
)

type TokenType int

// const (
	// itemError TokenType = iota
	// itemEOF
// )

type Token struct {
	// typ TokenType
	// val string
	R rune
	Val string
}

type StateFn func(*Lexer) StateFn

type Lexer struct {
	tokenCh chan Token
	src     *bufio.Reader
	startState StateFn
}

const EofRune = -1

func NewLexer(ch chan Token, src *bufio.Reader, start StateFn) *Lexer {
	l := &Lexer{
		tokenCh: ch,
		src: src,
		startState: start,
	}
	return l
}

func (l *Lexer) Next() rune {
	r, _, err := l.src.ReadRune()
	if err  != nil {
		// EOF or error
		return  EofRune
	}
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

