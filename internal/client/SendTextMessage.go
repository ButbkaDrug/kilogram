package client

import (
	"fmt"
	"log"

	tdlib "github.com/zelenin/go-tdlib/client"
    . "github.com/butbkadrug/kilogram/internal/models"
)

// Sends a text message to a chat specified as dest
// prints Chat ID and Message ID of the newly created message
func SendTextMessage(dest int64, text string) {

    kg := GetChats(true)

    kg.Waitgroup.Add(1)

    l := kg.Tdlib.GetListener()

    go handleSendMessageUpdates(kg, l)

    formattedText := &tdlib.FormattedText{
        Text: text,
    }

    content := &tdlib.InputMessageText{
        Text: formattedText,
        DisableWebPagePreview: true,
        ClearDraft: true,
    }

    if dest == 0 {

        me, err := kg.Tdlib.GetMe()

        if err != nil {
            log.Fatal("Destination id is not provided. Can't get self id ether, aborting: ", err)
        }

        dest = me.Id
    }


    _, err := kg.Tdlib.SendMessage(&tdlib.SendMessageRequest{
        ChatId: dest,
        InputMessageContent: content,
        ReplyMarkup: nil,
        ReplyTo: nil,
        Options: nil,
        MessageThreadId: 0,


    })

    if err != nil {
        log.Fatal("Failed to send the message: ", err)
    }

    kg.Waitgroup.Wait()
}

func handleSendMessageUpdates(kg *Kilogram, l *tdlib.Listener) {
    defer kg.Waitgroup.Done()

    for update := range l.Updates {

        switch u := update.(type) {
        case *tdlib.UpdateNewChat:
            kg.Chats[u.Chat.Id] = u.Chat
        case *tdlib.UpdateMessageSendSucceeded:
            fmt.Printf("%d %d", u.Message.ChatId, u.Message.Id)
            return
        case *tdlib.UpdateMessageSendFailed:
            fmt.Println("Failed to send the message!")
            return
        }
    }
}
