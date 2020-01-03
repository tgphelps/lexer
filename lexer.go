package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type tokenType int

const (
	itemError tokenType = iota
	itemEOF
)

type token struct {
	// typ tokenType
	// val string
	r rune
	val string
}

type lexer struct {
	tokenCh chan token
	src     *bufio.Reader
}

func delay(n time.Duration) {
	time.Sleep(n * time.Millisecond)
}

func newLexer(ch chan token, src *bufio.Reader) *lexer {
	l := &lexer{
		tokenCh: ch,
		src: src,
	}
	return l
}

func (l *lexer) run() {
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
		l.tokenCh <- token{r: r, val: string(r)}
	}
	close(l.tokenCh)
}

func main() {
	ch := make(chan token)
	src := bufio.NewReader(os.Stdin)
	l := newLexer(ch, src)
	fmt.Println("lexer:", l)

	go l.run()
	delay(100)
	for {
		tok, ok := <-ch
		fmt.Println("channel data:", tok, "ok:", ok)
		delay(100)
		if !ok {
			fmt.Println("got NOT ok:", ok)
			break
		}
	}
}
