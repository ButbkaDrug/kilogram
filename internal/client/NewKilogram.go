package client

import(
    "sync"
    . "github.com/butbkadrug/kilogram/internal/models"
    tdlib "github.com/zelenin/go-tdlib/client"
)
func NewKilogram () *Kilogram {
        td := NewClient()
    return &Kilogram{
        Tdlib: td,
        Waitgroup: &sync.WaitGroup{},
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
