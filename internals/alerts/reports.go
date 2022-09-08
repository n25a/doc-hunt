package alerts

import "github.com/fatih/color"

var R []report

type report struct {
	path    string
	line    int
	code    string
	Message string
}

func AddReport(path string, line int, code string, Message string) {
	R = append(R, report{
		path:    path,
		line:    line,
		code:    code,
		Message: Message,
	})
}

func GetReports() {
	color.Blue("============================================")
	color.Cyan("                   Reports                  ")
	color.Blue("============================================")

	if len(R) == 0 {
		color.Green("No reports. Your code is awesome!")
		return
	}

	for _, r := range R {
		color.Yellow("Path: %s", r.path)
		color.Yellow("Line: %d", r.line)
		color.Yellow("Code: %s", r.code)
		Warning.Printf("Message: %s\n", r.Message)
		color.Blue("============================================")
	}
}
