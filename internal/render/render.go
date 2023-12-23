package render

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/butbkadrug/kilogram/internal/models"
	tdlib "github.com/zelenin/go-tdlib/client"
    lg "github.com/charmbracelet/lipgloss"
)


const (
    FirstCollumnWidth = 15
    SecondCollumnWidth = 50
)

var(
)

//Function that pretty prints selected chat's messages
func PrintChat(chat *Kilochat){
    const w = 60
    var(
        t []string

        headerStyle = lg.NewStyle().
            Width(w).
            Padding(0,1).
            Foreground(lg.Color(color9)).
            Background(lg.Color(color0))

        replyStyle = lg.NewStyle().
            Width(w).
            Foreground(lg.Color(color13)).
            Italic(true).
            Align(lg.Right)
        bodyStyle = lg.NewStyle().
            Width(w).
            Padding(0, 0, 1, 1).
            Foreground(lg.Color(color9))
    )


    for i:=len(chat.Positions)-1; i>=0; i-- {
        var head, body, reply string
        id := chat.Positions[i]
        msg := chat.Messages[id]
        head = strconv.Itoa(int(id))

        if msg.IsOutgoing  {
            headerStyle.Align(lg.Right)
            head = "Me : " + head
        }

        head = headerStyle.Render(head)


        if msg.ReplyTo != nil {
            switch r := msg.ReplyTo.(type) {
            case *tdlib.MessageReplyToMessage:
                reply = getMessageText(chat.Messages[r.MessageId])
            default:
                reply = "Unknown message..."
            }
            reply = replyStyle.Render(reply)
        }




        body = getMessageText(msg)
        body = bodyStyle.Render(body)
        message := lg.JoinVertical(
            lg.Left,
            head,
            // reply,
            body,
        )


        t = append(t, message)

        if headerStyle.GetAlign() == lg.Right { headerStyle.Align(lg.Left) }
    }
    text := strings.Join(t, "\n")

    fmt.Print(text)
}


func getMessageText(msg *tdlib.Message) string {
    if msg == nil { return "" }
    switch m := msg.Content.(type){
    case *tdlib.MessageText:
        return m.Text.Text
    case *tdlib.MessagePhoto:
        return m.Caption.Text
    default:
        return msg.Content.MessageContentType()
    }
}
