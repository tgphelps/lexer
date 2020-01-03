package lexer

import (
	"bufio"
	"fmt"
)

type TokenType int

const (
	itemError TokenType = iota
	itemEOF
)

type Token struct {
	// typ TokenType
	// val string
	r rune
	val string
}

type Lexer struct {
	tokenCh chan Token
	src     *bufio.Reader
}

type StateFn func(*Lexer) StateFn

func NewLexer(ch chan Token, src *bufio.Reader) *Lexer {
	l := &Lexer{
		tokenCh: ch,
		src: src,
	}
	return l
}

func (l *Lexer) Run() {
	fmt.Println("lexer running...")

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
	close(l.tokenCh)
}

