
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

func ScanPPLine() {
    var line int64 = 0
    IN.ForwardCUR++
    for IN.Buf[IN.ForwardCUR] == ' ' || IN.Buf[IN.ForwardCUR] == '\t' {
        IN.ForwardCUR++
    }

    // # 2 "hello.c" 2              on Linux
    if IsDigit(IN.Buf[IN.ForwardCUR]) {
        goto read_line
    }

    read_line:
        // line number
        for IsDigit(IN.Buf[IN.ForwardCUR]) {
            line = 10 * line + int64(IN.Buf[IN.ForwardCUR] - '0')
            IN.ForwardCUR++
        }
        TokenCoord.Ppline = line - 1

        // skip white space
        for IN.Buf[IN.ForwardCUR] == ' ' || IN.Buf[IN.ForwardCUR] == '\t' {
            IN.ForwardCUR++
        }

        // get the filenme in pp file: "hello.c " ----> hello.c
        IN.ForwardCUR++
        FileNameSize := 1
        FileNameStart := (int)(IN.ForwardCUR)
        for IN.Buf[IN.ForwardCUR] != '"' && !IS_EOF(IN.Buf[IN.ForwardCUR]) && IN.Buf[IN.ForwardCUR] != '\n' {
            IN.ForwardCUR++
            FileNameSize++
        }
        TokenCoord.FileName = string(IN.Buf[FileNameStart:FileNameStart+FileNameSize])
        for IN.Buf[IN.ForwardCUR] != '\n' && !IS_EOF(IN.Buf[IN.ForwardCUR]) {
            IN.ForwardCUR++
        }
}

func SkipWhiteSpace() {
    var ch byte

//again:
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
                ScanPPLine()
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


    // no ExtraWhiteSpace, so not necessary to repeat again
    //if (ExtraWhiteSpace != nil) {
    //    goto again
    //}
}

func ScanEscapeChar(wide int) byte {
    var v int = 0
    var overflow int = 0
    IN.ForwardCUR++
    switch (IN.Buf[IN.ForwardCUR]) {
        case 'a':
            return '\a'
        case 'b':
            return '\b'
        case 'f':
            return '\f'
        case 'n':
            return '\n'
        case 'r':
            return '\r'
        case 't':
            return '\t'
        case 'v':
            return '\v'
 
        case '\'':
        case '"':
        case '\\':
        case '?':
            return IN.Buf[IN.ForwardCUR - 1]
        case 'x':
            if (!IsHexDigit(IN.Buf[IN.ForwardCUR])) {
                fmt.Println("Expect hex digit")
                return 'x'
            }
            v = 0
            for IsHexDigit(IN.Buf[IN.ForwardCUR]) {
                /***************************************
                Bug?
                if(v >> (WCharType->size * 8-4 ))
                (1) WCharType->size == 2
                        0xABCD * 16 + value --> overflow
                        0x0ABC is OK.
                (2) WCharType->size == 4
                        0x12345678  * 16 + value --> overflow
                        0x01234567 is OK.
                ***************************************/
                if (v >> ((WCharType.Size)*8 - 4)) {
                    overflow = 1
                }
                // v= v * 16 + value,  value : 0-9  A-F
                if IsDigit(IN.Buf[IN.ForwardCUR]) {
                    v = (v << 4) + IN.Buf[IN.ForwardCUR] - '0'
                } else {
                    v = (v << 4) + strings.ToUpper(IN.Buf[IN.ForwardCUR]) - 'A' + 10
                }
                IN.ForwardCUR++
            }
            if (overflow || (! wide && v > 255)) {
                fmt.Println("Hexademical espace sequence overflow")
            }
            return v

        case '0':
        case '1':
        case '2':
        case '3':
        case '4':
        case '5':
        case '6':
        case '7': // \ddd octal
            v = IN.Buf[IN.ForwardCUR -1] - '0'
            if (IsOctDigit(IN.Buf[IN.ForwardCUR])) {
                IN.ForwardCUR++
                v = (v << 3) + IN.Buf[IN.ForwardCUR] - '0'
                if IsOctDigit(IN.Buf[IN.ForwardCUR]) {
                    IN.ForwardCUR++
                    v = (v << 3) + IN.Buf[IN.ForwardCUR] - '0'
                }
            }
            return v
        default:
            fmt.Println("Unrecognized escape sequence")
            return IN.Buf[IN.ForwardCUR]
    }
}
