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
	r, _, _ := l.src.ReadRune()
	fmt.Println("lexer read:", r)
	l.tokenCh <- token{r: r, val: "char"}
	close(l.tokenCh)
}

func main() {
	ch := make(chan token)
	src := bufio.NewReader(os.Stdin)
	l := newLexer(ch, src)
	fmt.Println("lexer:", l)

	go l.run()
	delay(100)
	tok, ok := <-ch
	fmt.Println("channel data:", tok, "ok:", ok)
	delay(100)
	tok, ok = <-ch
	fmt.Println("channel data:", tok, "ok:", ok)
}
