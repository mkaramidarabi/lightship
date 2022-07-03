package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   `lightship`,
	Short: `lightship is a request orchestrator`,
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	var commitHash, commitTime, isDirty string
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("could not read runtime build information")
	}
	for _, bs := range bi.Settings {
		switch bs.Key {
		case "vcs.revision":
			commitHash = bs.Value
		case "vcs.time":
			commitTime = bs.Value
		case "vcs.modified":
			isDirty = bs.Value
		}
	}
	rootCmd.Version = "0.0.0" + " commit: " + commitHash + " Date: " +
		commitTime + " Dirty: " + isDirty
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
