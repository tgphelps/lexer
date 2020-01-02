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
	typ tokenType
	val string
}

type lexer struct {
	tokenCh chan token
}

func newLexer(ch chan token) *lexer {
	l := &lexer{
		tokenCh: ch,
	}
	return l
}

func (l *lexer) run() {
	fmt.Println("lexer running...")
} 

func main() {
	ch := make(chan token)
	l := newLexer(ch)
	fmt.Println("lexer:", l)
	go l.run()
	time.Sleep(100 * time.Millisecond)
}
