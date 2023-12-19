package client

import (
	"fmt"
	"log"

	tdlib "github.com/zelenin/go-tdlib/client"
)

func GetMessage(cid, mid int64) {

    kg := GetChats(false)


    msg, err := kg.Tdlib.GetMessage(&tdlib.GetMessageRequest{
        ChatId: cid,
        MessageId: mid,
    })

    if err != nil {
        log.Fatal(err)
    }

    printMessage(msg)


}

func printMessage(msg *tdlib.Message) {

    var output string
    // fmt.Println(msg)

    switch m := msg.Content.(type) {
    case *tdlib.MessageText:

        output += m.Text.Text
    }

    fmt.Println(output)

}

