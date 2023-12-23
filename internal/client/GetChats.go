package client

import (
	"fmt"
	"os"
	"sort"

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

func sortChats(c map[int64]*tdlib.Chat, p map[tdlib.JsonInt64]int64) []*tdlib.Chat {
    var order []tdlib.JsonInt64
    var chats []*tdlib.Chat

    for pos := range p{

        order = append(order, pos)

    }
    sort.Slice(order, func(i, j int) bool {return i < j})

    for _, i := range order {

        id := p[i]

        chats = append(chats, c[id])

    }

    return chats
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
        }
    }

    kilogram.Waitgroup.Done()
}
