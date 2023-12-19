package models

import(
    tdlib "github.com/zelenin/go-tdlib/client"
)

type Kilochat struct{
    Messages map[int64]*tdlib.Message
    Positions []int64
}

func NewKilochat(size int) *Kilochat {

    return & Kilochat{
        Messages: make(map[int64]*tdlib.Message, size),
        Positions: make([]int64, size),
    }
}
