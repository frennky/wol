package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wol",
	Short: "Wake On LAN",
	Long:  `Wake On LAN cli tool for sending magic packet to target MAC.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
