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
        Positions: make([]int64, 100),
        Messages: make(map[int64]*tdlib.Message),
    }
}
