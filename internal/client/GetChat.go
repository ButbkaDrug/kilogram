package client

import (
	"log"

	tdlib "github.com/zelenin/go-tdlib/client"
    . "github.com/butbkadrug/kilogram/internal/models"
)



// TODO:
// 1. To make a Open Chat request when getting the chat. And make getting withou
//    viewing optinal





// Fetches messeges for the chat with specified id
// You can also specify limit to load nessesery number of messages(max 100)
func GetChat(id int64, limit int32) *Kilochat {

    var start int64

    kg := GetChats()
    chat := NewKilochat(int(limit))


    if id == 0 {

        me, err := kg.Tdlib.GetMe()

        if err != nil {
            log.Fatal("Can't load the chat: ", err)
        }

        id = me.Id
    }

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

        if len(messages.Messages) < 1 { break }


        for _, msg := range messages.Messages {


            if msg == nil { continue }

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
                log.Println("Coun't get Message reply: ", err)
            }


            chat.Messages[m.Id] = m
        }
    }

    return chat
}
