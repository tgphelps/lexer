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

var tokNames = [...]string{"err", "eof", "char", "alpha", "num", "sp"}

func main() {
	ch := make(chan lexer.Token)
	src := bufio.NewReader(os.Stdin)
	l := lexer.NewLexer(ch, src, lexState, tokNames[:])

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
		if l.GotEof {
			l.Emit(lexer.Token{Type: TokEof})
			break
		}
		r := l.Next()
		fmt.Printf("rune = %q\n", r)
		if r == lexer.EofRune {
			l.Emit(lexer.Token{Type: TokEof})
			break
		} else {
			var t lexer.Token
			l.UnNext()
			if unicode.IsLetter(r) {
				s := collect(l, unicode.IsLetter)
				t = lexer.Token{Type: TokAlphas, Val: s}
			} else if unicode.IsNumber(r) {
				s := collect(l, unicode.IsNumber)
				t = lexer.Token{Type: TokDigits, Val: s}
			} else if unicode.IsSpace(r) {
				s := collect(l, unicode.IsSpace)
				t = lexer.Token{Type: TokWhiteSpace, Val: s}
			} else {
				// re-fetch the character
				r := l.Next()
				t = lexer.Token{Type: TokChar, Val: string(r)}
			}
			l.Emit(t)
		}
	}
	return nil
}

func collect(l *lexer.Lexer, testFunc charTestFn) string {
	var b strings.Builder
	// b.WriteRune(r)
	for {
		r := l.Next()
		if testFunc(r) {
			b.WriteRune(r)
		} else {
			// We have hit the end of the string of chars we want
			// If we just got EOF, don't try to push it back
			if r != lexer.EofRune {
				l.UnNext()
			}
			break
		}
	}
	return b.String()
}

func delay(d time.Duration) {
	time.Sleep(d * time.Millisecond)
}
