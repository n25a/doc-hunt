package alerts

import "github.com/fatih/color"

var Error *color.Color
var Warning *color.Color

func init() {
	Error = color.New(color.FgRed, color.Bold)
	Warning = color.New(color.FgYellow, color.Bold)
}
