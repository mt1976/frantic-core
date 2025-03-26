package logHandler

import (
	"strings"

	"github.com/mt1976/frantic-core/colours"
)

func assembleLogFileName(in, name string) string {
	return in + name + ".log"
}

func formatNameWithColor(colour, name string) string {
	name = strings.ToUpper(name)
	return colour + "[" + name + "]" + Reset + " "
}

func setColoursNormal() {
	Reset = colours.Reset
	Red = colours.Red
	Green = colours.Green
	Yellow = colours.Yellow
	Blue = colours.Blue
	Magenta = colours.Magenta
	Cyan = colours.Cyan
	Gray = colours.Gray
	White = colours.White
}

func setColoursWindows() {
	Reset = ""
	Red = ""
	Green = ""
	Yellow = ""
	Blue = ""
	Magenta = ""
	Cyan = ""
	Gray = ""
	White = ""
}
