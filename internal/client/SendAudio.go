package client

import (
	"fmt"
	"os"

	tdlib "github.com/zelenin/go-tdlib/client"
)

type SendAudioParams struct {
    Id int64
    File, Caption string
}

func SendAudio(p *SendAudioParams){
    f, err := os.Stat(p.File)

    if err != nil {
        fmt.Fprintln(os.Stderr, "Can't open a file: ", err)
        os.Exit(1)
        return
    }

    inputFile := &tdlib.InputFileLocal{
        Path: p.File,

    }

    content := &tdlib.InputMessageAudio{
        Audio: inputFile,
        Title: f.Name(),
        Caption: &tdlib.FormattedText{
            Text: p.Caption,
        },
    }

    msg, err := SendMessage(&tdlib.SendMessageRequest{
        ChatId: p.Id,
        MessageThreadId: 0,
        ReplyTo: nil,
        Options: nil,
        ReplyMarkup: nil,
        InputMessageContent: content,
    })

    if err != nil {
        fmt.Fprintln(os.Stderr, "Can't send the message: ", err)
        os.Exit(1)
        return
    }

    fmt.Println(msg.ChatId, msg.Id)
}
