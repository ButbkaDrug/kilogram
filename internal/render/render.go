package render

import (
	"fmt"
	"strconv"

	. "github.com/butbkadrug/kilogram/internal/models"
	lg "github.com/charmbracelet/lipgloss"
	tdlib "github.com/zelenin/go-tdlib/client"
)
const w = 75

var(
    messageBodyStyle = lg.NewStyle().
        Width(w).
        Margin(0,1,1,0)
    messageHeaderStyle = lg.NewStyle().
        Width(w).
        Foreground(lg.Color("#d4d4d4")).
        Align(lg.Right)
    messageReplyStyle = lg.NewStyle().
        Italic(true).
        Width(w).
        Align(lg.Right).
        Foreground(lg.Color("#646464"))
    messageWrapperStyle = lg.NewStyle().
        Width(w)
)
//Function that pretty prints selected chat's messages
func PrintChat(chat *Kilochat){
    var t []string


    for i:=len(chat.Positions)-1; i>=0; i-- {
        var head, body, reply string
        id := chat.Positions[i]
        msg := chat.Messages[id]
        head = messageHeaderStyle.Render(strconv.Itoa(int(id)))

        if msg.ReplyTo != nil {
            switch r := msg.ReplyTo.(type) {
            case *tdlib.MessageReplyToMessage:
                reply = getMessageText(chat.Messages[r.MessageId])
            default:
                reply = "Unknown message..."
            }

            reply = messageReplyStyle.Render(reply)
        }



        body = messageBodyStyle.Render( getMessageText(msg))
        message := messageWrapperStyle.Render(
            lg.JoinVertical(lg.Left,
                head,
                reply,
                body,
            ),
        )
        t = append(t, message)
    }
    text := lg.JoinVertical(
        lg.Left,
        t...
    )

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
