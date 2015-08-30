package gubase

import (
    "encoding/json"
    "fmt"
    "os"
)

const PrefsFileName = ".gucli.json"

var prefs = make(map[string]interface{})

func GetPrefsPath() string {
    usr := GetCurrentUser()
    return fmt.Sprintf("%s/%s", usr.HomeDir, PrefsFileName)
}

func ReadPrefs() {
    prefsPath := GetPrefsPath()

    file, err := os.Open(prefsPath)
    defer file.Close()

    fmt.Fprintf(os.Stderr, "Reading preferences from %s\n", prefsPath)

    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    } else {
        dec := json.NewDecoder(file)
        dec.Decode(&prefs)
    }
}

func WritePrefs() {
    prefsPath := GetPrefsPath()

    file, err := os.Create(prefsPath)
    defer file.Close()

    fmt.Fprintf(os.Stderr, "Writing preferences to %s\n", prefsPath)

    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    } else {
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
