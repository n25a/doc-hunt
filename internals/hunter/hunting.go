package hunter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/n25a/doc-hunt/internals/alerts"
)

func Hunting(dir string, excludePath []string, excludeFile []string) error {
	// find all paths
	paths := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// remove files from excludeFile
	// remove paths from excludePath
	var tmp []string
	for _, path := range paths {
		add := true
		for _, exclude := range excludeFile {
			if strings.Contains(path, exclude) {
				add = false
				break
			}
		}
		for _, exclude := range excludePath {
			if strings.Contains(path, exclude) {
				add = false
				break
			}
		}

		if add {
			tmp = append(tmp, path)
		}
	}
	paths = tmp

	// check godoc for each file
	err = checkGodoc(paths)
	if err != nil {
		return err
	}

	// get report
	alerts.GetReports()

	return nil
}
