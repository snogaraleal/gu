package gubase

import (
    "encoding/json"
    "fmt"
    "os"
)

const PrefsPath = ".gucli.json"

var prefs = make(map[string]interface{})

func ReadPrefs() {
    file, err := os.Open(PrefsPath)
    defer file.Close()

    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    } else {
        fmt.Fprintf(os.Stderr, "Reading preferences from %s\n", PrefsPath)

        dec := json.NewDecoder(file)
        dec.Decode(&prefs)
    }
}

func WritePrefs() {
    file, err := os.Create(PrefsPath)
    defer file.Close()

    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    } else {
        fmt.Fprintf(os.Stderr, "Writing preferences to %s\n", PrefsPath)

        enc := json.NewEncoder(file)
        enc.Encode(&prefs)
    }
}

func GetPref(name string) interface{} {
    value, _ := prefs[name]
    return value
}

func SetPref(name string, value interface{}) {
    prefs[name] = value
}
