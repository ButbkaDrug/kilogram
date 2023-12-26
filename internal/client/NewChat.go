package client

import (
	"fmt"
	"os"

	tdlib "github.com/zelenin/go-tdlib/client"
)

func NewChat(id int64) *tdlib.Chat {
    kg := GetChats()

    chat, err := kg.Tdlib.CreatePrivateChat(&tdlib.CreatePrivateChatRequest{
        UserId: id,
        Force: false,
    })

    if err != nil {
        fmt.Fprintf(os.Stderr, "Can't create a chat: %s", err)
    }

    return chat
}
