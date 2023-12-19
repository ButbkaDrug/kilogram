package render

import(
    "github.com/butbkadrug/kilogram/internal/client"
    "github.com/charmbracelet/lipgloss"
)

//Function that pretty prints selected chat's messages
func PrintChat(chat *Kilochat){


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
