package main

import (
	"bufio"
	"fmt"
	"lexer"
	"os"
	"time"
)

// type stateFn func(*lexer.Lexer) stateFn

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
        fmt.Println("channel data:", tok, "ok:", ok)
        delay(100)
        if !ok {
            fmt.Println("got NOT ok:", ok)
            break
        }
    }
}

func lexState(l *lexer.Lexer) lexer.StateFn {
	fmt.Println("start lexState")
	r :=  l.Next()
	t := lexer.Token{R: r, Val: string(r)}
	l.Emit(t)
	return nil
}

func delay(d time.Duration) {
	time.Sleep(d * time.Millisecond)
}
