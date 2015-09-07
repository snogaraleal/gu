package gusp

import (
    "flag"
    "fmt"
    "os"

    "golang.org/x/crypto/ssh/terminal"

    "github.com/snogaraleal/gu/gubase"
)

func Do(action string, args []string) {
    switch action {
    case "auth":
        DoAuth(args)
    case "syllabus":
        DoSyllabus(args)
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

func DoSyllabus(args []string) {
    // Get params
    var flagSet = flag.NewFlagSet("syllabus", flag.ExitOnError)

    var query string
    flagSet.StringVar(&query, "q", "", "Search query")

    flagSet.Parse(args)

    // Make request
    session := NewSession()
    result := session.Request(&SyllabusUtility{query}).(SyllabusResult)

    // Show results
    for _, course := range result.Courses {
        // Description
        fmt.Printf(
            gubase.Color("● %s » %s\n", gubase.ColorWhite),
            course.Code, course.Title)

        // Documents
        for _, doc := range course.Docs {
            fmt.Printf("  ● %s\n", doc.Title)
            fmt.Printf("    %s\n", doc.Link)
        }

        fmt.Println("")
    }

    fmt.Println(gubase.Color(result.Message, gubase.ColorWhite))
}
