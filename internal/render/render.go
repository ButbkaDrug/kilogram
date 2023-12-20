package render

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/butbkadrug/kilogram/internal/models"
	tdlib "github.com/zelenin/go-tdlib/client"
)
const w = 75

//Function that pretty prints selected chat's messages
func PrintChat(chat *Kilochat){
    var t []string


    for i:=len(chat.Positions)-1; i>=0; i-- {
        var head, body, reply string
        id := chat.Positions[i]
        msg := chat.Messages[id]
        head = strconv.Itoa(int(id))

        if msg.IsOutgoing  {
            head = "Me >> " + head
        }

        if msg.ReplyTo != nil {
            switch r := msg.ReplyTo.(type) {
            case *tdlib.MessageReplyToMessage:
                reply = getMessageText(chat.Messages[r.MessageId])
            default:
                reply = "Unknown message..."
            }
        }



        body = getMessageText(msg)
        message := head + reply + body
        t = append(t, message)
    }
    text := strings.Join(t, "\n")

    fmt.Print(text)
}

func getMessageText(msg *tdlib.Message) string {
    switch m := msg.Content.(type){
    case *tdlib.MessageText:
        return m.Text.Text
    case *tdlib.MessagePhoto:
        return m.Caption.Text
    default:
        return msg.Content.MessageContentType()
    }
}
        // text += "|--------------------------------------------------|\n"

        // if msg.ReplyTo != nil {

        //     text +="|\t\t ->"

        //     switch r := msg.ReplyTo.(type) {
        //     case *tdlib.MessageReplyToMessage:
        //         text += getMessageText(msgs[r.MessageId])
        //     default:
        //         text += "Unknown message..."
        //     }

        //     text += "\n"
        // }

        // text += fmt.Sprintf("| %d |\t%s\n", msg.Id, getMessageText(msg))

    // }
    // text += "|--------------------------------------------------|\n"
