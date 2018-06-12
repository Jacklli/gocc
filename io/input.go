
package input

import (
    "os"
)

type coord struct {
    filename string
    ppline int
    line int
    col int
}

type input struct {
    filename string
    base *byte
    cursor *byte
    lineHead *byte
    line int
    fp *os.File
    size uint32
}

func (in *input) readSourceFile(filename *string) *byte {

    return in.base
}

func (in *input) closeSourceFile() {
    in.fp.Close()
}
