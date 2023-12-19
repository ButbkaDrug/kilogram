package models

import(
    tdlib "github.com/zelenin/go-tdlib/client"
)

type Kilochat struct{
    Messages map[int64]*tdlib.Message
    Positions []int64
}
