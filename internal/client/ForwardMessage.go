package client

import (
	"fmt"
	"log"

	. "github.com/butbkadrug/kilogram/internal/models"
	tdlib "github.com/zelenin/go-tdlib/client"
)

// TODO:
// 1. I want to return new message, to make it possible to chain this command
//    with get chat, or get message. And any other message in fact

func ForwardMessage(source, dest int64, messages []int64) {

    var output []any

    kg := GetChats(true)

    lsr := kg.Tdlib.GetListener()

    defer lsr.Close()

    kg.Waitgroup.Add(1)

    go updateHandler(kg, lsr)

    if dest == 0 {
        me, err := kg.Tdlib.GetMe()

        if err != nil {
            log.Fatal("Destenation chat id not provided. Forwardig to saved messages faild too. aborting...")
        }
        dest = me.Id
    }


    if len(messages) < 1 {
        chat, err := kg.Tdlib.GetChat(&tdlib.GetChatRequest{ChatId: source})

        if err != nil {
            log.Fatalf("Forwarding faild! ERROR: %s", err)
        }

        source = chat.Id

        lastMessage := chat.LastMessage
        if lastMessage == nil {
            log.Fatal("Message id not provided! No last message found ether! Aborting...")
        }

        messages = append(messages, lastMessage.Id)
    }

    fmsgs, err := kg.Tdlib.ForwardMessages(&tdlib.ForwardMessagesRequest{
        ChatId: dest,
        FromChatId: source,
        MessageIds: messages,
        SendCopy: false,
        Options: nil,
        OnlyPreview: false,
        RemoveCaption: false,
        MessageThreadId: 0,
    })

    if err != nil {
        log.Fatal("Can't forward messages: ", err)
    }

    kg.Waitgroup.Wait()

    output = append(output, dest)

    for _, msg := range fmsgs.Messages {
        output = append(output, msg.Id)
    }

    fmt.Print(output...)
}

func updateHandler(kg *Kilogram, l *tdlib.Listener) {
    fmt.Printf("Waiting for the message to be forwarder...")
    defer kg.Waitgroup.Done()
    for update := range l.Updates {
        switch u := update.(type){
        case *tdlib.UpdateMessageSendSucceeded:
            return
        case *tdlib.UpdateMessageSendFailed:
            fmt.Printf("Failed to forward the message! Error code: %d\n", u.ErrorCode)
            return
        case *tdlib.UpdateDeleteMessages:
            fmt.Print("Delete Message update...!\n")
            return
        }
    }
}
