package color

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
	"testing"
)

func TestColor(_ *testing.T) {
	fmt.Println("NoColor:", NoColor, isatty.IsTerminal(os.Stdout.Fd()))
	s := String("cat", FgRed)
	fmt.Println(s)
}
