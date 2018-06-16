
package lex

import (
    "github.com/gocc/io"
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

func ScanPPLine() {
    line := 0
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
            line = 10 * line + int(IN.Buf[IN.ForwardCUR]) - '0'
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
