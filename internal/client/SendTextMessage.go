package client

import (
	tdlib "github.com/zelenin/go-tdlib/client"
)

// Sends a text message to a chat specified as dest
// prints Chat ID and Message ID of the newly created message
func SendTextMessage(dest int64, text string) (*tdlib.Message, error){

    formattedText := &tdlib.FormattedText{
        Text: text,
    }

    content := &tdlib.InputMessageText{
        Text: formattedText,
        DisableWebPagePreview: true,
        ClearDraft: true,
    }

    return SendMessage(&tdlib.SendMessageRequest{
        ChatId: dest,
        InputMessageContent: content,
        ReplyMarkup: nil,
        ReplyTo: nil,
        Options: nil,
        MessageThreadId: 0,
    })
}
