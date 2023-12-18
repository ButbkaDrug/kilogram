package client

import (
	"fmt"
	"log"

	tdlib "github.com/zelenin/go-tdlib/client"
)

func ForwardMessage(source, dest int64, messages []int64) {

    kg := GetChats(true)

    lsr := kg.tdlib.GetListener()

    defer lsr.Close()

    kg.waitgroup.Add(1)

    go updateHandler(kg, lsr)

    if dest == 0 {
        me, err := kg.tdlib.GetMe()

        if err != nil {
            log.Fatal("Destenation chat id not provided. Forwardig to saved messages faild too. aborting...")
        }
        dest = me.Id
    }


    if len(messages) < 1 {
        chat, err := kg.tdlib.GetChat(&tdlib.GetChatRequest{ChatId: source})

        if err != nil {
            log.Fatalf("Forwarding faild! ERROR: %s", err)
        }

        source = chat.Id

        lastMessage := chat.LastMessage
        if lastMessage == nil {
            log.Fatal("Message id not provided! No last message found ether")
        }

        messages = append(messages, lastMessage.Id)
    }

    fmsgs, err := kg.tdlib.ForwardMessages(&tdlib.ForwardMessagesRequest{
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

    kg.waitgroup.Wait()

    fmt.Printf("Forwared %d message(s)", fmsgs.TotalCount)
}

func updateHandler(kg *Kilogram, l *tdlib.Listener) {
    fmt.Printf("Waiting for the message to be forwarder...")
    defer kg.waitgroup.Done()
    for update := range l.Updates {
        switch u := update.(type){
        case *tdlib.UpdateMessageSendSucceeded:
            fmt.Print("Message successfully forwarded!\n")
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
