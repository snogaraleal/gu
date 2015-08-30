package gubase

import (
    "fmt"
    "os"

    "golang.org/x/crypto/ssh/terminal"
)

const (
    ColorBlack = 0
    ColorRed = 1
    ColorGreen = 2
    ColorYellow = 3
    ColorBlue = 4
    ColorMagenta = 5
    ColorCyan = 6
    ColorWhite = 7
)

func Color(text string, color int) string {
    if terminal.IsTerminal(int(os.Stdout.Fd())) {
        return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", 30 + color, text)
    } else {
        return text
    }
}
