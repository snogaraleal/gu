package main

import (
    "fmt"
    "os"

    "github.com/snogaraleal/gu/gubase"
    "github.com/snogaraleal/gu/gusgs"
)

type Command struct {
    Title string
    Do func(string, []string)
}

var Registry = map[string]Command {
    "sgs": {"SGS Student Housing", gusgs.Do},
}

func main() {
    notice()

    if len(os.Args) < 3 {
        usage()
        os.Exit(1)
    }

    command, action := os.Args[1], os.Args[2]
    target, ok := Registry[command]

    if ok {
        fmt.Fprintln(os.Stderr, gubase.Color(target.Title, gubase.ColorWhite))
        target.Do(action, os.Args[3:])
    } else {
        fmt.Fprintln(
            os.Stderr,
            gubase.Color("Command not found", gubase.ColorRed))
        os.Exit(1)
    }
}

func usage() {
    fmt.Fprintln(os.Stderr, gubase.Color("** Usage **", gubase.ColorWhite))
    fmt.Fprintln(os.Stderr, "gucli <command> <action>")
}

func notice() {
    fmt.Fprintln(os.Stderr, gubase.Color("** GU-CLI **", gubase.ColorWhite))
    fmt.Fprintln(os.Stderr, "")

    fmt.Fprintln(
        os.Stderr,
        "This software is a collection of utilities that may be useful\n" +
        "to you when studying at the University of Gothenburg.\n")

    fmt.Fprintln(
        os.Stderr,
        "However, this software has *no official status* and\n" +
        "*no affiliation* or whatsoever with the University of Gothenburg.\n")
}
