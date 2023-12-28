package render

import (
	"fmt"

	tdlib "github.com/zelenin/go-tdlib/client"
    lg "github.com/charmbracelet/lipgloss"
)

type RenderUsersParams struct {
    Verbose bool
    All bool
    Limit int
    Offset int
}

func RenderUsers(users []*tdlib.User, params *RenderUsersParams) {
    var results string


    for _, user := range users {
        if _, ok := user.Status.(*tdlib.UserStatusOnline); !params.All && !ok {
            continue
        }

        results += VerboseRender(user) + "\n"

    }

    fmt.Print(results)
}

func VerboseRender(user *tdlib.User) string{
    var name, id, phone, status string

    var nameStyle = lg.NewStyle().Bold(true).Width(30)
    var phoneStyle = lg.NewStyle().Foreground(lg.Color(color15)).Width(15)
    var idStyle = lg.NewStyle().Bold(true).Background(lg.Color(color4))
    var statusStyle = lg.NewStyle().Foreground(lg.Color(color10))

    name = formatName(user)
    name = nameStyle.Render(name)

    id = fmt.Sprintf("%d ", user.Id)
    id = idStyle.Render(id)


    phone = phoneStyle.Render(user.PhoneNumber)

    status = formatStatus(user)
    if status != "Online" { statusStyle.Foreground(lg.Color(color11)) }
    status = statusStyle.Render(status)

    line := "-----------------------------------------------------------"

    return lg.JoinHorizontal(
        lg.Left,
        name,
        phone,
        id,
        status,
    )+ "\n" + line

}
