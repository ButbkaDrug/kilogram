package render

import (
	"fmt"
	"time"

	// lg "github.com/charmbracelet/lipgloss"
	tdlib "github.com/zelenin/go-tdlib/client"
)

type RenderUserParams struct {
    Verbose  bool
    Status bool
    Username bool
    Phone bool
}

func RenderUser(user *tdlib.User, params *RenderUserParams){
    var output string

    if params.Username || params.Verbose {
        output += formatName(user) + "\n"
    }

    if params.Phone || params.Verbose {
        output += user.PhoneNumber + "\n"
    }

    if params.Status || params.Verbose {
        output += formatStatus(user) + "\n"
    }

    fmt.Println(output)
}

func formatStatus(user *tdlib.User) string {
    var output string

    switch status := user.Status.(type){
    case *tdlib.UserStatusOnline:
        output = "Online"
    case *tdlib.UserStatusEmpty:
        output = "Empty"
    case *tdlib.UserStatusOffline:
        t := time.Unix(int64(status.WasOnline), 0)
        elapsed := time.Since(t)
        output = formatElapsedTime(elapsed)
    case *tdlib.UserStatusRecently:
        output = "Recently"
    case *tdlib.UserStatusLastWeek:
        output = "Last week"
    case *tdlib.UserStatusLastMonth:
        output = "Last month"

    }

    return output
}

func formatName(user *tdlib.User) string {
    var output string

    output = fmt.Sprintf("%s %s", user.FirstName, user.LastName)

    if output != "" { return output }

    if len(user.Usernames.ActiveUsernames) > 0 {
        output = user.Usernames.ActiveUsernames[0]
    }

    if output != "" { return output }

    output = fmt.Sprintf("%d", user.Id)
    if !user.HaveAccess {
        output += "(Not Accessable)"
    }

    return output

}

func formatElapsedTime(t time.Duration) string {
    var output string

    d := int(t.Hours()) / 24
    h := int(t.Hours()) % 24
    m := int(t.Minutes()) % 60
    s := int(t.Seconds()) % 60

    if d > 0 {
        output += fmt.Sprintf("%d d ", d)
    }

    if h > 0 {
        output += fmt.Sprintf("%d h ", h)
    }

    if m > 0 && d < 0{
        output += fmt.Sprintf("%d m ", m)
    }

    if h < 0  && d < 0{
        output += fmt.Sprintf("%d s", s)
    }


    return output

}
