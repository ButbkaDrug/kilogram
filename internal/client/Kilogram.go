package client

import (
    "sync"
	tdlib "github.com/zelenin/go-tdlib/client"
)

type Kilogram struct {
    tdlib *tdlib.Client
    waitgroup *sync.WaitGroup
    Chats map[int64]*tdlib.Chat
    Positions map[tdlib.JsonInt64]int64
    Users map[int64]*tdlib.User
    Messages map[int64]*tdlib.Message
}

func NewKilogram () *Kilogram {
        td := NewClient()
    return &Kilogram{
        tdlib: td,
        waitgroup: &sync.WaitGroup{},
        Chats: make(map[int64]*tdlib.Chat, 100),
        Users: make(map[int64]*tdlib.User, 100),
        Positions: make(map[tdlib.JsonInt64]int64),
        Messages: make(map[int64]*tdlib.Message),
    }
}

func getMessageText(msg *tdlib.Message) string {
    switch m := msg.Content.(type){
    case *tdlib.MessageText:
        return m.Text.Text
    case *tdlib.MessagePhoto:
        return m.Caption.Text
    default:
        return msg.Content.MessageContentType()
    }
}
