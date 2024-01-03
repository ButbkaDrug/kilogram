package client

import (
	"fmt"
	"os"

	tdlib "github.com/zelenin/go-tdlib/client"
)

type SendDocumentParams struct {
    ChatId int64
    Files []string
    Caption string
}

func SendDocument(p *SendDocumentParams) (*tdlib.Messages, error){

    var err error
    var content []tdlib.InputMessageContent

    for _, file := range p.Files {
        if _, err = os.Stat(file); err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }

        inputFile := &tdlib.InputFileLocal{Path: file}

        c := &tdlib.FormattedText{
            Text: p.Caption,
            Entities: nil,
        }

        content = append(content, &tdlib.InputMessageDocument{
            Document: inputFile,
            Caption: c,
            Thumbnail: nil,
            DisableContentTypeDetection: false,
        })
    }

    return SendAlbum(&tdlib.SendMessageAlbumRequest{
        ChatId: p.ChatId,
        OnlyPreview: false,
        MessageThreadId: 0,
        ReplyTo: nil,
        Options: nil,
        InputMessageContents: content,
    })
}
