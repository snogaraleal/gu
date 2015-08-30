package gusgs

import (
    "flag"
    "fmt"

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
    /*
     * Get params
     */
    var flagSet = flag.NewFlagSet("search", flag.ExitOnError)

    var marketName string
    flagSet.StringVar(&marketName, "m", "", fmt.Sprintln(DefaultMarkets))

    flagSet.Parse(args)

    var market int
    if len(marketName) > 0 {
        market = DefaultMarkets[marketName]
    }

    /*
     * Make request
     */
    result := Request(&SearchUtility{market}).(SearchResult)

    /*
     * Show results
     */
    fmt.Println("")

    for _, item := range result.Items {
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

        fmt.Printf("\n")
    }

    fmt.Printf(
        gubase.Color("%d results\n", gubase.ColorWhite),
        result.TotalCount)
}

func DoAuth(args []string) {
    /*
     * Ask credentials
     */
}
