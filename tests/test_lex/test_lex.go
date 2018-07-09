
package main

import (
    "fmt"
    "github.com/gocc/lex"
)

func main() {
    for i:=0; i<lex.TK_END; i++ {
        fmt.Println(lex.TokenStrings[i])
    }
}
