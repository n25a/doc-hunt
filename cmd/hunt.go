package cmd

import (
	"github.com/n25a/doc-hunt/internals/alerts"
	"github.com/n25a/doc-hunt/internals/hunter"
	"github.com/n25a/doc-hunt/internals/parser"
	"github.com/spf13/cobra"
)

var (
	dir         string
	excludePath string
	excludeFile string
)

var huntCMD = &cobra.Command{
	Use:   "hunt",
	Short: "Search in directory for checking godoc",
	Run:   hunt,
}

func init() {
	huntCMD.Flags().StringVarP(&dir, "dir", "d", "", "Project directory")
	huntCMD.Flags().StringVar(&excludePath, "exclude-path", "", "Project directory")
	huntCMD.Flags().StringVar(&excludeFile, "exclude-file", "", "Project directory")
}

func hunt(_ *cobra.Command, _ []string) {
	if dir == "" {
		_, _ = alerts.Error.Println("Error: dir is not set")
		panic("Error: dir is not set")
	}

	var newExcludePath []string
	var newExcludeFile []string
	if excludePath != "" {
		newExcludePath = parser.ExportData(excludePath)
	}
	if excludeFile != "" {
		newExcludeFile = parser.ExportData(excludeFile)
	}

	err := hunter.Hunting(dir, newExcludePath, newExcludeFile)
	if err != nil {
		_, _ = alerts.Error.Println(err)
		panic(err)
	}
}
