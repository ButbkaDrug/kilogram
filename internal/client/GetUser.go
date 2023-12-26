package client

import (
	"log"

	tdlib "github.com/zelenin/go-tdlib/client"
)

func GetUser(id int64) *tdlib.User {
    kg := GetChats()

    user, err := kg.Tdlib.GetUser(&tdlib.GetUserRequest{UserId: id})

    if err != nil {
        log.Fatal("Can't fetch the user info: ", err)
    }

    return user

}
