package gusgs

import (
    "flag"
    "fmt"
    "os"

    "golang.org/x/crypto/ssh/terminal"
    "github.com/snogaraleal/gu/gubase"
)

func Do(action string, args []string) {
    switch action {
    case "search":
        DoSearch(args)
    case "auth":
        DoAuth(args)
    }
}

func DoSearch(args []string) {
    // Get params
    var flagSet = flag.NewFlagSet("search", flag.ExitOnError)

    var marketName string
    flagSet.StringVar(&marketName, "m", "", fmt.Sprintln(DefaultMarkets))

    flagSet.Parse(args)

    var market int
    if len(marketName) > 0 {
        market = DefaultMarkets[marketName]
    }

    // Make request
    result := Request(&SearchUtility{market}).(SearchResult)

    // Handle result
    for _, item := range result.Items {
        fmt.Println("")

        // Description
        fmt.Printf(
            gubase.Color("● %s » %s\n", gubase.ColorWhite),
            item.SeekArea, item.Address)
        fmt.Printf(
            "%s, %s m², floor %s\n",
            item.Description, item.Area, item.Floor)

        // Rent
        fmt.Printf(gubase.Color("%.0f SEK\n", gubase.ColorGreen), item.Rent)

        // Properties
        for _, property := range item.Properties {
            fmt.Printf("  ● %s\n", property.Description)
        }

        // Dates
        fmt.Printf("        Apply before: %s\n", item.LastDay)
        fmt.Printf("           Free from: %s\n", item.FreeFrom)
    }

    fmt.Printf(
        gubase.Color("\n%d results\n", gubase.ColorWhite),
        result.TotalCount)
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
    result := Request(&AuthUtility{username, password}).(AuthResult)

    // Handle result
    if result.Success {
        fmt.Println("Successfully authenticated")

        // Store token
        gubase.ReadPrefs()
        gubase.SetPref("sgs.token", result.Token)
        gubase.WritePrefs()
    } else {
        fmt.Println("Authentication failure")
    }
}
