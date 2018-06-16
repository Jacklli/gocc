package io

import (
    "os"
    "github.com/gocc/utils"
)

type Coord struct {
    FileName string
    Ppline int
    Line int
    Col int
}

type Input struct {
    FileName string
    Buf []byte
    Size int64
    LineHead int64
    Line int64
    ForwardCUR int64
    Fp *os.File
}

func (in *Input) ReadSourceFile(fileName string) {
    fp, err := os.Open(fileName)
    utils.CheckErr(err)
    defer fp.Close()
    in.Fp = fp
 
    size, err := fp.Seek(0, os.SEEK_END)
    in.Size = size
    fp.Seek(0, os.SEEK_SET)
    in.Buf = make([]byte, size)
    fp.Read(in.Buf)
    utils.CheckErr(err)
}

