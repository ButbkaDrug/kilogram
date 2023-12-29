package client

import (
	"fmt"
	"os"
	. "github.com/butbkadrug/kilogram/internal/models"
	tdlib "github.com/zelenin/go-tdlib/client"
)

// Sends a LoadChats request to tdlib.
// Returns a pointer to the Kilogram instance
func GetChats() *Kilogram {


    kg := NewKilogram()

    listener := kg.Tdlib.GetListener()
    kg.Waitgroup.Add(1)
    go handleUpdates(kg, listener)


    defer listener.Close()

    r := &tdlib.LoadChatsRequest{
        ChatList: &tdlib.ChatListMain{},
        Limit: 100,
    }

    var i = 0

    for {
        _, err := kg.Tdlib.LoadChats(r)
        i++
        if err != nil {
            break
        }
    }

    var j = 0

    for {
        ok, err := kg.Tdlib.LoadChats(&tdlib.LoadChatsRequest{
            ChatList: &tdlib.ChatListArchive{},
            Limit: 1,
        })
        if ok != nil {
            fmt.Println("\n\n", ok.Type, ok.Extra, "\n\n")
        }

        j++

        if err != nil {
            break
        }
    }
    fmt.Println(i, " requested main list")
    fmt.Println(j, " requested archived list")



    kg.Waitgroup.Wait()


    return kg
}


func handleUpdates(kilogram *Kilogram, l *tdlib.Listener) {

    // loop:
    for update := range l.Updates {
        fmt.Printf("%v\n", update)
        switch upd := update.(type) {
        case *tdlib.Ok:
            fmt.Printf("%#v\n", upd)
            // break loop
        case *tdlib.UpdateUser:
            // kilogram.Users[upd.User.Id] = upd.User
        case *tdlib.UpdateNewChat:
            kilogram.Chats[upd.Chat.Id] = upd.Chat
        case *tdlib.UpdateChatPosition:
            // do something
            kilogram.Positions = append(kilogram.Positions, upd.ChatId)
        case *tdlib.UpdateSupergroup:
            // groups = append(groups, upd.Supergroup)
        case *tdlib.UpdateChatReadInbox:
            kilogram.Chats[upd.ChatId].UnreadCount = upd.UnreadCount
        case *tdlib.UpdateChatLastMessage:
            if _, ok := kilogram.Chats[upd.ChatId]; !ok {
                fmt.Fprintf(os.Stderr, "Chat not found! %v", upd)
                continue
            }
            kilogram.Chats[upd.ChatId].LastMessage = upd.LastMessage
        default:
        }
    }

    kilogram.Waitgroup.Done()
}
