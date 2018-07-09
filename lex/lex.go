
package lex

import (
    "github.com/gocc/io"
    "fmt"
)

var (
    TokenValue = TK_value {}
    END_OF_FILE byte = 255

    PeekPoint *byte
    PeekValue = TK_value {}
    PeekCoord = TK_value {}

    IN = io.Input {}
    TokenCoord = io.Coord {}
    PrevCoord = io.Coord {}
)

func IsDigit(c byte) bool {
    return c >= '0' && c <= '9'
}

func IsOctDigit(c byte) bool {
    return c >= '0' && c <= '7'
}

func IsHexDigit(c byte) bool {
    return IsDigit(c) || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f')
}

func IsLetter(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c == '_') || (c >= 'A' && c <= 'Z')
}

func IsLetterOrDigit(c byte) bool {
    return IsLetter(c) || IsDigit(c)
}

func IS_EOF(c byte) bool {
    return c == END_OF_FILE && IN.ForwardCUR == IN.Size
}

func GetNextToken() int {
    var tok int
    PrevCoord = TokenCoord
    SkipWhiteSpace()
    TokenCoord.Line = IN.Line
    TokenCoord.Col  = (int64)(IN.ForwardCUR - IN.LineHead + 1)
    // use function pointer table to avoid a large switch statement.
    // tok = (*Scanners[IN.Buf[ForwardCur]])()
    return tok
}

func SkipWhiteSpace() {
    var ch byte
    ch = IN.Buf[IN.ForwardCUR]
    for ch == '\t' || ch == '\v' || ch == '\f' || ch == ' ' ||
        ch == '\r' || ch == '\n' || ch == '/'  || ch == '#' {
        switch ch {
            case '\n':
                TokenCoord.Ppline++
                IN.Line++
                IN.ForwardCUR++
                IN.LineHead = IN.ForwardCUR
                break

            case '#':
                break

            case '/':           // comments
                if IN.Buf[IN.ForwardCUR+1] != '/' && IN.Buf[IN.ForwardCUR+1] != '*' {
                    return
                }
                IN.ForwardCUR++
                if IN.Buf[IN.ForwardCUR] == '/' {
                    IN.ForwardCUR++
                    for IN.Buf[IN.ForwardCUR] != '\n' && !IS_EOF(IN.Buf[IN.ForwardCUR]) {
                        IN.ForwardCUR++
                    }
                } else {
                    IN.ForwardCUR++
                    for IN.Buf[IN.ForwardCUR] != '*' || IN.Buf[IN.ForwardCUR+1] != '/' {
                        if IN.Buf[IN.ForwardCUR] == '\n' {
                            TokenCoord.Ppline++
                            IN.Line++
                        } else if (IS_EOF(IN.Buf[IN.ForwardCUR]) || IS_EOF(IN.Buf[IN.ForwardCUR+1])) {
                            fmt.Println("Comment is not closed.")
                            return
                        }
                        IN.ForwardCUR++
                    }
                    IN.ForwardCUR += 2
                }
                break
                
            default:
                IN.ForwardCUR++
                break
        }
        ch = IN.Buf[IN.ForwardCUR]
    }
}

