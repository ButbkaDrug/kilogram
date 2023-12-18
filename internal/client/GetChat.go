package client

import (
	"fmt"
	"log"

	tdlib "github.com/zelenin/go-tdlib/client"
)

func GetChat(id int64) {
    kg := GetChats(true)

    // Fetch messages for chat with given id

    var (
        start int64
        limit int32
        msgs map[int64]*tdlib.Message
        positions []int64
    )

    limit = 10
    start = 0
    msgs = make(map[int64]*tdlib.Message, limit)

    for {
        messages, err := kg.tdlib.GetChatHistory(&tdlib.GetChatHistoryRequest{
        ChatId: id,
        Limit: limit,
        Offset: 0,
        OnlyLocal: false,
        FromMessageId: start,
        })

        if err != nil {
            log.Fatal(err)
        }

        for _, msg := range messages.Messages {
            msgs[msg.Id] = msg
            positions = append(positions, msg.Id)
        }

        start += messages.Messages[len(messages.Messages)-1].Id
        limit -= messages.TotalCount

        if limit <= 1 { break }
    }

    for _, msg := range msgs {
        reply := msg.ReplyTo
        if reply == nil { continue }
        switch r := reply.(type){
        case *tdlib.MessageReplyToMessage:

            if _, ok := msgs[r.MessageId]; ok { continue }
            m, err := kg.tdlib.GetMessage(&tdlib.GetMessageRequest{
                ChatId: r.ChatId,
                MessageId: r.MessageId,
            })

            if err != nil {
                log.Println("Coun't get Message reply...")
            }

            msgs[m.Id] = m
        }
    }


    printChatHistory(msgs, positions)
}

func printChatHistory(msgs map[int64]*tdlib.Message, positions []int64) {


    var text string
    for i:=len(positions)-1; i>=0; i-- {
        id := positions[i]
        msg := msgs[id]

        text += "|--------------------------------------------------|\n"

        if msg.ReplyTo != nil {

            text +="|\t\t ->"

            switch r := msg.ReplyTo.(type) {
            case *tdlib.MessageReplyToMessage:
                text += getMessageText(msgs[r.MessageId])
            default:
                text += "Unknown message..."
            }

            text += "\n"
        }

        text += fmt.Sprintf("| %d |\t%s\n", msg.Id, getMessageText(msg))

    }
    text += "|--------------------------------------------------|\n"
    fmt.Println(text)
}

func(k *Kilogram) getMessage(chatId, messageId int64) *tdlib.Message {

    m, err := k.tdlib.GetMessage(&tdlib.GetMessageRequest{
        ChatId: chatId,
        MessageId: messageId,
    })

    if err != nil {
        log.Fatal("Can't get the reply message")
    }

    return m
}
