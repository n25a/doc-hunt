package hunter

import (
	"bufio"
	"os"
	"strings"

	"github.com/n25a/doc-hunt/internal/alerts"
)

func checkGodoc(paths []string) error {
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(f)

		comments := []string{}
		lineCounter := 0
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)
			lineCounter += 1

			//if strings.Contains(line, "//") {
			//	comments = append(comments, line)
			//}

			if strings.Contains(line, "package") {
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Package without godoc comments",
					)
				} else {
					comments = []string{}
				}
			}

			if strings.Contains(line, "func") {
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Function without godoc comments",
					)
				} else {
					comments = []string{}
				}
			}

			if strings.Contains(line, "const") {
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Constance without godoc comments",
					)
				} else {
					comments = []string{}
				}
			}

			if strings.Contains(line, "struct") {
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Struct without godoc comments",
					)
				} else {
					comments = []string{}
				}
			}

			if strings.Contains(line, "var") {
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Global Variable without godoc comments",
					)
				} else {
					comments = []string{}
				}
			}

		}
	}

	return nil
}
