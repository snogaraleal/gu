package gusp

import (
    "fmt"

    "github.com/snogaraleal/gu/gubase"
)

func Do(action string, args []string) {
    switch action {
    case "auth":
        DoAuth(args)
    }
}

func DoAuth(args []string) {
    fmt.Println(gubase.Color("SP", gubase.ColorWhite))
}
