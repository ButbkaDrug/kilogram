package client

import (
	"os"

	tdlib "github.com/zelenin/go-tdlib/client"
)

type SendPhotoParams struct {
    ChatId int64
    Path string
    Width, Height int32
    Caption string
    Spoiler bool
}

func SendPhoto(p *SendPhotoParams) (*tdlib.Message, error) {

    var msg *tdlib.Message
    var err error

    _, err = os.Stat(p.Path)

    if err != nil {
        return msg, err
    }

    inputFile := &tdlib.InputFileLocal{Path: p.Path}

    c := &tdlib.FormattedText{
        Text: p.Caption,
        Entities: nil,
    }

    content := &tdlib.InputMessagePhoto{
        Photo: inputFile,
        Thumbnail: nil,
        AddedStickerFileIds: nil,
        Width: p.Width,
        Height: p.Height,
        Caption: c,
        HasSpoiler: p.Spoiler,
    }

    return SendMessage(&tdlib.SendMessageRequest{
        ChatId: p.ChatId,
        ReplyTo: nil,
        Options: nil,
        MessageThreadId: 0,
        ReplyMarkup: nil,
        InputMessageContent: content,
    })
}
