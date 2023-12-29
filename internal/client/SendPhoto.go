package client

import (
	"fmt"
	"os"

	tdlib "github.com/zelenin/go-tdlib/client"
)

type SendPhotoParams struct {
    ChatId int64
    Files []string
    Width, Height int32
    Caption string
    Spoiler bool
}

func SendPhoto(p *SendPhotoParams) (*tdlib.Messages, error){

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

        content = append(content, &tdlib.InputMessagePhoto{
            Photo: inputFile,
            Thumbnail: nil,
            AddedStickerFileIds: nil,
            Width: p.Width,
            Height: p.Height,
            Caption: c,
            HasSpoiler: p.Spoiler,
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
