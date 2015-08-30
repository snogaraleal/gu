package gubase

import (
    "os/user"
)

func GetCurrentUser() *user.User {
    usr, err := user.Current()

    if err != nil {
        panic("Cannot obtain current user")
    }

    return usr
}
