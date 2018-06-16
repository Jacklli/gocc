
package main

import (
    "fmt"
    "github.com/gocc/io"
)

func main() {
    in := io.Input{}
    in.ReadSourceFile("test.txt")
    fmt.Printf("%d\n", in.Size)
    var i int64 
    fmt.Printf("%s\n", string(in.Buf))
    for i=0; i<in.Size; i++ {
        fmt.Printf("%s\n", string(in.Buf[i]))
    }
}
