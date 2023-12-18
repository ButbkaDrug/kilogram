package client

import (
	"fmt"
	"log"

	tdlib "github.com/zelenin/go-tdlib/client"
)

func SendTextMessage(dest int64, text string) {

    kg := GetChats(true)

    kg.waitgroup.Add(1)

    l := kg.tdlib.GetListener()

    go handleSendMessageUpdates(kg, l)

    formattedText := &tdlib.FormattedText{
        Text: text,
    }

    content := &tdlib.InputMessageText{
        Text: formattedText,
        DisableWebPagePreview: true,
        ClearDraft: true,
    }


    _, err := kg.tdlib.SendMessage(&tdlib.SendMessageRequest{
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

    kg.waitgroup.Wait()
}

func handleSendMessageUpdates(kg *Kilogram, l *tdlib.Listener) {
    defer kg.waitgroup.Done()

    for update := range l.Updates {

        switch u := update.(type) {
        case *tdlib.UpdateNewChat:
            kg.Chats[u.Chat.Id] = u.Chat
        case *tdlib.UpdateMessageSendSucceeded:
            fmt.Println("Message successfully sent!")
            return
        case *tdlib.UpdateMessageSendFailed:
            fmt.Println("Failed to send the message!")
            return
        }
    }
}
