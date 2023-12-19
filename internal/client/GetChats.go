package client

import (
	"fmt"
	tdlib "github.com/zelenin/go-tdlib/client"
    . "github.com/butbkadrug/kilogram/internal/models"
)


// Sends a LoadChats request to tdlib.
// Returns a pointer to the Kilogram instance
func GetChats(all bool) *Kilogram {

    kg := NewKilogram()

    listener := kg.Tdlib.GetListener()
    kg.Waitgroup.Add(1)
    go handleUpdates(kg, listener)


    defer listener.Close()

    r := &tdlib.LoadChatsRequest{
        ChatList: &tdlib.ChatListMain{},
        Limit: 100,
    }

    for {
        _, err := kg.Tdlib.LoadChats(r)
        if err != nil {
            break
        }
    }

    for {
        _, err := kg.Tdlib.LoadChats(&tdlib.LoadChatsRequest{
            ChatList: &tdlib.ChatListArchive{},
            Limit: 1,
        })

        if err != nil {
            break
        }
    }


    kg.Waitgroup.Wait()

    return kg
}

func handleUpdates(kilogram *Kilogram, l *tdlib.Listener) {

    loop:
    for update := range l.Updates {
        switch upd := update.(type) {
        case *tdlib.Ok:
            break loop
        case *tdlib.UpdateUser:
            // kilogram.Users[upd.User.Id] = upd.User
        case *tdlib.UpdateNewChat:
            kilogram.Chats[upd.Chat.Id] = upd.Chat
        case *tdlib.UpdateChatPosition:
            // do something
            kilogram.Positions[upd.Position.Order] = upd.ChatId
        case *tdlib.UpdateSupergroup:
            // groups = append(groups, upd.Supergroup)
        case *tdlib.UpdateChatReadInbox:
            kilogram.Chats[upd.ChatId].UnreadCount = upd.UnreadCount
        case *tdlib.UpdateChatLastMessage:
            kilogram.Chats[upd.ChatId].LastMessage = upd.LastMessage
        }
    }

    kilogram.Waitgroup.Done()
}

func PrintChats(chats map[int64]*tdlib.Chat, all bool) {

    for id, chat := range chats {

        if !all && chat.UnreadCount < 1 { continue }

        fmt.Printf("%d\t%s\t%d\t", id, chat.Title, chat.UnreadCount)

        if chat.LastMessage == nil { continue }

        fmt.Printf("%d\n", chat.LastMessage.Id)
    }

}
