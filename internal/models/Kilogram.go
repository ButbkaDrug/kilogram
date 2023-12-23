package models

import (
    "sync"
	tdlib "github.com/zelenin/go-tdlib/client"
)

type Kilogram struct {
    Tdlib *tdlib.Client
    Waitgroup *sync.WaitGroup
    Chats map[int64]*tdlib.Chat
    Positions []int64
    Users map[int64]*tdlib.User
    Messages map[int64]*tdlib.Message
}

