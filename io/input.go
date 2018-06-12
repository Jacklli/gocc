package io

import (
    "os"
    "github.com/gocc/utils"
)

type Coord struct {
    fileName string
    ppline int
    line int
    col int
}

type Input struct {
    fileName string
    base *byte
    cursor *byte
    lineHead *byte
    line int
    fp *os.File
    Size int64
}

func (in *Input) ReadSourceFile(fileName string) *byte {
    fp, err := os.Open(fileName)
    utils.CheckErr(err)
    defer fp.Close()
    in.fp = fp
 
    size, err := fp.Seek(0, os.SEEK_END)
    in.Size = size
    fp.Seek(0, os.SEEK_SET)
    buf := make([]byte, size)
    fp.Read(buf)
    utils.CheckErr(err)
    in.base = &buf[0]

    return in.base
}
