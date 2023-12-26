package client

import (
	"fmt"
	"log"
	"os"

	tdlib "github.com/zelenin/go-tdlib/client"
)

func GetUsers() []*tdlib.User {

    var results []*tdlib.User

    kg := NewKilogram()

    users, err := kg.Tdlib.GetContacts()

    if err != nil {
        log.Fatal("Can't load contacts: ", err)
    }

    for _, id := range users.UserIds {

        user, err := kg.Tdlib.GetUser(&tdlib.GetUserRequest{UserId: id})

        if err != nil {
            fmt.Fprintln(os.Stderr, "Faild to laod the user: ", err)
            continue
        }

        results = append(results, user)
    }

    return results
}
