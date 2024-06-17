package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var cmd = &cobra.Command{Use: "ytdl"}

func Start() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cmd.AddCommand(apiCmd)
}
