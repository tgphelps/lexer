package main

import (
	"fmt"
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
	n int
	val string
}

type lexer struct {
	tokenCh chan token
}

func delay(n time.Duration) {
	time.Sleep(n * time.Millisecond)
}

func newLexer(ch chan token) *lexer {
	l := &lexer{
		tokenCh: ch,
	}
	return l
}

func (l *lexer) run() {
	fmt.Println("lexer running...")
	l.tokenCh <- token{n: 1, val: "one"}
	close(l.tokenCh)
}

func main() {
	ch := make(chan token)
	l := newLexer(ch)
	fmt.Println("lexer:", l)
	go l.run()
	delay(100)
	tok, ok := <-ch
	fmt.Println("channel data:", tok, "ok:", ok)
	delay(100)
	tok, ok = <-ch
	fmt.Println("channel data:", tok, "ok:", ok)
	delay(100)
	tok, ok = <-ch
	fmt.Println("channel data:", tok, "ok:", ok)
}
