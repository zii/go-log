package color

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
	"strings"
)

var (
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. This is a global option and affects all colors. For more control
	// over each color block use the methods DisableColor() individually.
	NoColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))
)

const escape = "\x1b"

// SGR Code

// Base attributes
const (
	Reset        = "0"
	Bold         = "1"
	Faint        = "2"
	Italic       = "3"
	Underline    = "4"
	BlinkSlow    = "5"
	BlinkRapid   = "6"
	ReverseVideo = "7"
	Concealed    = "8"
	CrossedOut   = "9"
)

// Foreground text colors
const (
	FgBlack   = "30"
	FgRed     = "31"
	FgGreen   = "32"
	FgYellow  = "33"
	FgBlue    = "34"
	FgMagenta = "35"
	FgCyan    = "36"
	FgWhite   = "37"
)

func format(attrs ...string) string {
	seq := strings.Join(attrs, ";")
	return fmt.Sprintf("%s[%sm", escape, seq)
}

func unformat() string {
	return fmt.Sprintf("%s[%sm", escape, Reset)
}

func String(s string, attrs ...string) string {
	if NoColor {
		return s
	}

	return format(attrs...) + s + unformat()
}
