package main

import (
	"bufio"
	"fmt"
	"lexer"
	"os"
	"time"
)

// type stateFn func(*lexer.Lexer) stateFn

var mytoken lexer.TokenType = 1

func main() {
	fmt.Println("hello", mytoken)
    ch := make(chan lexer.Token)
    src := bufio.NewReader(os.Stdin)
    l := lexer.NewLexer(ch, src)
    fmt.Println("lexer:", l)

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

func delay(d time.Duration) {
	time.Sleep(d * time.Millisecond)
}
