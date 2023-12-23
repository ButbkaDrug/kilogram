package client

import (
	"fmt"
	"log"
	"os"

	. "github.com/butbkadrug/kilogram/internal/models"
	tdlib "github.com/zelenin/go-tdlib/client"
)

// TODO:
// Implement LIMIT functionality

type ForwardMessageParams struct {
    // Chat id the messages belong to
    Source int64
    // Destination chat id, where you want to send your messages
    Dest int64
    // Ids of the messages to be forwarded.
    // If no IDs provided, last message found in the chat will be forwarded
    Messages []int64
    // If no id provided, number of messages to forward starting from the last
    // message found in the chat
    Limit int32
}

func ForwardMessage(p *ForwardMessageParams) {

    var output []any

    kg := GetChats()

    lsr := kg.Tdlib.GetListener()

    defer lsr.Close()

    kg.Waitgroup.Add(1)

    go updateHandler(kg, lsr)

    if p.Dest == 0 {
        me, err := kg.Tdlib.GetMe()

        if err != nil {
            log.Fatal("Destenation chat id not provided. Forwardig to saved messages faild too. aborting...")
        }
        p.Dest = me.Id
    }


    if len(p.Messages) < 1 {
        chat, err := kg.Tdlib.GetChat(&tdlib.GetChatRequest{ChatId: p.Source})

        if err != nil {
            log.Fatalf("Forwarding faild! ERROR: %s", err)
        }

        p.Source = chat.Id

        lastMessage := chat.LastMessage
        if lastMessage == nil {
            log.Fatal("Message id not provided! No last message found ether! Aborting...")
        }

        p.Messages = append(p.Messages, lastMessage.Id)
    }

    fmsgs, err := kg.Tdlib.ForwardMessages(&tdlib.ForwardMessagesRequest{
        ChatId: p.Dest,
        FromChatId: p.Source,
        MessageIds: p.Messages,
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

    output = append(output, p.Dest)

    for _, msg := range fmsgs.Messages {
        if msg == nil { continue }
        output = append(output, msg.Id)
    }

    fmt.Print(output...)
}

func updateHandler(kg *Kilogram, l *tdlib.Listener) {
    defer kg.Waitgroup.Done()
    for update := range l.Updates {
        switch u := update.(type){
        case *tdlib.UpdateMessageSendSucceeded:
            return
        case *tdlib.UpdateMessageSendFailed:
            fmt.Fprintf(os.Stderr, "Failed to forward the message! Error code: %d\n", u.ErrorCode)
            return
        case *tdlib.UpdateDeleteMessages:
            fmt.Fprint(os.Stderr, "Delete Message update...!\n")
            return
        }
    }
}
