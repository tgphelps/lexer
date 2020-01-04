package main

import (
	"bufio"
	"fmt"
	"lexer"
	"os"
	"strings"
	"time"
	"unicode"
)

// type stateFn func(*lexer.Lexer) stateFn
type charTestFn func(rune) bool

// The first two of these must match those in lexer.go
const (
	TokError lexer.TokenType = iota
	TokEof
	TokChar
	TokAlphas
	TokDigits
	TokWhiteSpace
)

func main() {
	ch := make(chan lexer.Token)
	src := bufio.NewReader(os.Stdin)
	l := lexer.NewLexer(ch, src, lexState)

	go l.Run()
	delay(100)
	for {
		tok, ok := <-ch
		if !ok {
			// fmt.Println("got NOT ok:", ok)
			break
		}
		fmt.Println("channel data:", tok, "ok:", ok)
		delay(100)
	}
}

func lexState(l *lexer.Lexer) lexer.StateFn {
	fmt.Println("start lexState")
	for {
		r := l.Next()
		// fmt.Printf("rune = %q\n", r)
		if r == lexer.EofRune {
			l.Emit(lexer.Token{Type: TokEof})
			break
		} else {
			var t lexer.Token
			if unicode.IsLetter(r) {
				s := collect(l, unicode.IsLetter, r)
				t = lexer.Token{Type: TokAlphas, Val: s}
			} else if unicode.IsNumber(r) {
				s := collect(l, unicode.IsNumber, r)
				t = lexer.Token{Type: TokDigits, Val: s}
			} else if unicode.IsSpace(r) {
				s := collect(l, unicode.IsSpace, r)
				t = lexer.Token{Type: TokWhiteSpace, Val: s}
			} else {
				t = lexer.Token{Type: TokChar, Val: string(r)}
			}
			l.Emit(t)
		}
	}
	return nil
}

func collect(l *lexer.Lexer, testFunc charTestFn, r rune) string {
	var b strings.Builder
	b.WriteRune(r)
	for {
		r := l.Next()
		if testFunc(r) {
			b.WriteRune(r)
		} else {
			l.UnNext()
			break
		}
	}
	return b.String()
}

func delay(d time.Duration) {
	time.Sleep(d * time.Millisecond)
}
