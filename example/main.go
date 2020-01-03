package main

import (
	"bufio"
	"fmt"
	"lexer"
	"os"
	"time"
)

// type stateFn func(*lexer.Lexer) stateFn

// The first two of these must match those in lexer.go
const (
	TokError lexer.TokenType = iota
	TokEof
	TokChar
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
		r :=  l.Next()
		if r == lexer.EofRune {
			l.Emit(lexer.Token{Type: TokEof})
			break
		} else {
			t := lexer.Token{Type: TokChar, Val: string(r)}
			l.Emit(t)
		}
	}
	return nil
}

func delay(d time.Duration) {
	time.Sleep(d * time.Millisecond)
}
