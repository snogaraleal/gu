package gusp

import (
    "fmt"
    "os"

    "golang.org/x/crypto/ssh/terminal"

    "github.com/snogaraleal/gu/gubase"
)

func Do(action string, args []string) {
    switch action {
    case "auth":
        DoAuth(args)
    }
}

func DoAuth(args []string) {
    // Get params
    var username, password string

    fmt.Printf("Username: ")
    fmt.Scanf("%s", &username)

    fmt.Printf("Password: ")
    raw, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
    password = string(raw)
    fmt.Println("")

    // Make request
    session := NewSession()
    result := session.Request(&AuthUtility{username, password}).(AuthResult)
    result.SyncPrefs()

    // Show result
    if result.Success {
        fmt.Println(gubase.Color("● Auth succeeded", gubase.ColorGreen))
    } else {
        fmt.Println(gubase.Color("● Auth failed", gubase.ColorRed))
    }
}
