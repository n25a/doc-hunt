package hunter

import (
	"bufio"
	"os"
	"strings"

	"github.com/n25a/doc-hunt/internals/alerts"
)

func checkGodoc(paths []string) error {
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(f)

		var (
			comments       = []string{}
			lineCounter    = 0
			funcFind       = false
			constFind      = false
			structFind     = false
			varFind        = false
			anotherBracket = 0
		)

		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)
			lineCounter += 1

			if strings.Contains(line, "{") {
				anotherBracket++
			}
			if strings.Contains(line, "}") {
				anotherBracket--
			}

			if funcFind && anotherBracket > 0 {
				continue
			}

			if !funcFind && !constFind && !structFind && strings.Contains(line, "//") {
				comments = append(comments, line)
			}

			if strings.Contains(line, "package") {
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Package without godoc comments",
					)
				} else {
					comments = []string{}
				}
			}

			if strings.Contains(line, "func") && strings.Contains(line, "{") {
				funcFind = true
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Function without godoc comments",
					)
				} else {
					comments = []string{}
				}
			}

			if !funcFind && strings.Contains(line, "const") && (strings.Contains(line, "(") ||
				strings.Contains(line, "=")) {
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Constance without godoc comments",
					)
				} else {
					comments = []string{}
				}

				if strings.Contains(line, "const (") {
					constFind = true
				}
			}

			if !funcFind && strings.Contains(line, "struct {") {
				structFind = true
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Struct without godoc comments",
					)
				} else {
					comments = []string{}
				}
			}

			if !funcFind && strings.Contains(line, "var ") && (strings.Contains(line, "(") ||
				strings.Contains(line, "=")) {
				if len(comments) == 0 {
					alerts.AddReport(
						path, lineCounter, line, "Global Variable without godoc comments",
					)
				} else {
					comments = []string{}
				}
				if strings.Contains(line, "var (") {
					varFind = true
				}
			}

			if funcFind && anotherBracket == 0 && strings.Contains(line, "}") {
				funcFind = false
			}

			if constFind && strings.Contains(line, ")") {
				constFind = false
			}

			if structFind && strings.Contains(line, "}") {
				structFind = false
			}

			if varFind && strings.Contains(line, ")") {
				varFind = false
			}

			if len(comments) > 0 && !strings.Contains(line, "//") {
				comments = []string{}
			}
		}
	}

	return nil
}
