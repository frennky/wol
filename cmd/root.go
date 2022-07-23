package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "wol",
	Long: "Simple Wake On LAN CLI tool.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
