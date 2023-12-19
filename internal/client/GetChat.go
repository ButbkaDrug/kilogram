package client

import (
	"log"

	tdlib "github.com/zelenin/go-tdlib/client"
    . "github.com/butbkadrug/kilogram/internal/models"
)





// Fetches messeges for the chat with specified id
// You can also specify limit to load nessesery number of messages(max 100)
func GetChat(id int64, limit int32) *Kilochat {

    var start int64

    kg := GetChats(true)
    chat := NewKilochat(int(limit))

    for {
        messages, err := kg.Tdlib.GetChatHistory(&tdlib.GetChatHistoryRequest{
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
            m, err := kg.Tdlib.GetMessage(&tdlib.GetMessageRequest{
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
