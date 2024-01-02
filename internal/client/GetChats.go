package client

import (
	"fmt"
	"os"
	. "github.com/butbkadrug/kilogram/internal/models"
	tdlib "github.com/zelenin/go-tdlib/client"
)

type state struct {
    requests int
    responses int
    done bool
}


// Sends a LoadChats request to tdlib.
// Returns a pointer to the Kilogram instance
func GetChats() *Kilogram {

    state := &state{
        requests: 0,
        responses: 0,
        done: false,
    }


    kg := NewKilogram()

    listener := kg.Tdlib.GetListener()
    kg.Waitgroup.Add(1)
    go handleUpdates(kg, listener, state)


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

        state.requests += 1
    }


    for {
        _, err := kg.Tdlib.LoadChats(&tdlib.LoadChatsRequest{
            ChatList: &tdlib.ChatListArchive{},
            Limit: 100,
        })

        if err != nil {
            state.done = true
            break
        }
        state.requests += 1
    }


    kg.Waitgroup.Wait()


    return kg
}


func handleUpdates(kilogram *Kilogram, l *tdlib.Listener, s *state) {

    for update := range l.Updates {
        if s.done && s.requests == s.responses { break }
        switch upd := update.(type) {
        case *tdlib.Ok:
            s.responses += 1
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
