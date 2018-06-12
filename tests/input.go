
package main

import (
    "fmt"
    "github.com/gocc/io"
)

func main() {
    in := io.Input{}
    buf := in.ReadSourceFile("test.txt")
    fmt.Printf("%d\n", in.Size)
    for i:=0; i<in.Size; i++ {
        fmt.Printf("%s\n", string(buf[i]))
    }
}
