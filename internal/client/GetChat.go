package client

import (
	"fmt"
	"log"

	tdlib "github.com/zelenin/go-tdlib/client"
)

type Kilochat struct{
    Messages map[int64]*tdlib.Message
    Positions []int64
}

func NewKilochat(size int) *Kilochat {

    return & Kilochat{
        Messages: make(map[int64]*tdlib.Message, size),
        Positions: make([]int64, size),
    }
}



// Fetches messeges for the chat with specified id
// You can also specify limit to load nessesery number of messages(max 100)
func GetChat(id int64, limit int32) *Kilochat {

    var start int64

    kg := GetChats(true)
    chat := NewKilochat(int(limit))

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
            chat.Messages[msg.Id] = msg
            chat.Positions = append(chat.Positions, msg.Id)
        }

        start += messages.Messages[len(messages.Messages)-1].Id
        limit -= messages.TotalCount

        if limit < 1 { break }
    }

    for _, msg := range chat.Messages{
        reply := msg.ReplyTo
        if reply == nil { continue }
        switch r := reply.(type){
        case *tdlib.MessageReplyToMessage:

            if _, ok := chat.Messages[r.MessageId]; ok { continue }
            m, err := kg.tdlib.GetMessage(&tdlib.GetMessageRequest{
                ChatId: r.ChatId,
                MessageId: r.MessageId,
            })

            if err != nil {
                log.Println("Coun't get Message reply...")
            }

            chat.Messages[m.Id] = m
        }
    }

    return chat
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
