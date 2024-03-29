package client

import (
	"fmt"
	"os"

	. "github.com/butbkadrug/kilogram/internal/models"
	tdlib "github.com/zelenin/go-tdlib/client"
)

func SendMessage(r *tdlib.SendMessageRequest) (*tdlib.Message, error){

    kg := GetChats()

    if r.ChatId == 0 {

        me, err := kg.Tdlib.GetMe()

        if err != nil {
            fmt.Fprintln(os.Stderr, "Chat id not provided. Failed to forward to saved messages too: ", err)
            os.Exit(1)
        }

        r.ChatId = me.Id

    }


    l := kg.Tdlib.GetListener()
    kg.Waitgroup.Add(1)
    defer kg.Waitgroup.Wait()

    go func(kg *Kilogram, l *tdlib.Listener) {
        defer kg.Waitgroup.Done()

        for update := range l.Updates {

            switch u := update.(type) {
            case *tdlib.UpdateNewChat:
                kg.Chats[u.Chat.Id] = u.Chat
            case *tdlib.UpdateMessageSendSucceeded:
                return
            case *tdlib.UpdateMessageSendFailed:
                fmt.Fprintln(os.Stderr, "Failed to send the message!")
                return
            case *tdlib.Error:
                return
            case *tdlib.UpdateFile:
                //
            }
        }

    }(kg, l)


    msg, err := kg.Tdlib.SendMessage(r)

    if err != nil {
        return nil, err
    }

    return msg, nil
}
