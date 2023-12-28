package render

import (
	"fmt"
	"strings"

	lg "github.com/charmbracelet/lipgloss"
	tdlib "github.com/zelenin/go-tdlib/client"
)


type RenderChatsParams struct {
    Limit int
    Offset int
    PrintAll bool
    Verbose bool
    Order []int64
}

const (
    color0 = "#1C1C1C"
    color1 = "#AF5F5F"
    color2 = "#5F875F"
    color3 = "#87875F"
    color4 = "#5F87AF"
    color5 = "#5F5F87"
    color6 = "#5F8787"
    color7 = "#6C6C6C"
    color8 = "#444444"
    color9 = "#FF8700"
    color10 = "#87AF87"
    color11 = "#FFFFAF"
    color12 = "#8FAFD7"
    color13 = "#8787AF"
    color14 = "#5FAFAF"
    color15 = "#FFFFFF"
)

func RenderChats (chats map[int64]*tdlib.Chat, params *RenderChatsParams) {

    var chatList []*tdlib.Chat

    for _, id := range params.Order {
        chat := chats[id]

        if chat == nil { continue }

        chatList = append(chatList, chat)
    }

    chatList = filterChats(chatList, params)


    if params.Verbose {
        renderChatsVerbose(chatList)
        return
    }

    renderChatsMinified(chatList)
}

func renderChatsVerbose(chats []*tdlib.Chat){
    const (
        firstColWidth = 11
        secondColWidth = 25
        thirdColWidth = 4
        forthColWidth = 25
    )
    var(

        output string

        firstColStyle = lg.NewStyle().
            Width(firstColWidth).
            Foreground(lg.Color(color5))

        secondColStyle = lg.NewStyle().
            Width(secondColWidth).
            Foreground(lg.Color(color10))
        thirdColStyle = lg.NewStyle().
            Width(thirdColWidth).
            Foreground(lg.Color(color1)).
            Align(lg.Center)
        forthColStyle = lg.NewStyle().
            Width(forthColWidth).
            Foreground(lg.Color(color9))

    )


    for _, chat := range chats {
        var(
            row string
            id string
            title string
            counter string
            message string
        )

        id = shrinkTextLine(fmt.Sprint(chat.Id), firstColWidth)
        id = firstColStyle.Render(id)


        title = shrinkTextLine(chat.Title, secondColWidth)
        title = secondColStyle.Render(title)


        counter = fmt.Sprint(chat.UnreadCount)
        counter = thirdColStyle.Render(counter)


        message = getMessageText(chat.LastMessage)
        message = shrinkTextLine(message, forthColWidth)
        message = forthColStyle.Render(message)

        row += lg.JoinHorizontal(
            lg.Left,
            id,
            title,
            counter,
            message,
        )
        output += row + "\n"
    }

    fmt.Print(output)

}


func renderChatsMinified(chats []*tdlib.Chat) {
    var output string

    for _, chat := range chats {
        output += fmt.Sprintln(chat.Id)
    }

    if len(output) < 1 { return }

    fmt.Println(output)
}

func filterChats(chats []*tdlib.Chat, params *RenderChatsParams) []*tdlib.Chat {
    var result []*tdlib.Chat

    for i, chat := range chats {

        if !params.PrintAll && chat.UnreadCount < 1 { continue }
        if params.Offset > 0 && i < params.Offset { continue }
        if params.Limit > 0 && len(result) >= params.Limit { break }

        result = append(result, chat)
    }

    return result
}

func shrinkTextLine(s string, w int) string {
    if w < 4 || len(s) <= w {
        return s
    }

    s = strings.ReplaceAll(s, "\n", "->")

    return s[:w]
}
