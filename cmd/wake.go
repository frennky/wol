package cmd

import (
	"log"

	"github.com/frennky/wol/pkg/types"
	"github.com/frennky/wol/pkg/wol"
	"github.com/spf13/cobra"
)

var mac string
var host string
var port int

var wakeCmd = &cobra.Command{
	Use:  "wake",
	Long: "Sends magic packet to target MAC.",
	Run:  runWakeCmd,
}

func init() {
	rootCmd.AddCommand(wakeCmd)
	wakeCmd.Flags().StringVarP(&mac, "mac", "", "", "Target MAC Address (required)")
	_ = wakeCmd.MarkFlagRequired("mac")
	wakeCmd.Flags().StringVarP(&host, "host", "", wol.DefaultBroadcastIPAddress, "Broadcast IP Address (optional)")
	wakeCmd.Flags().IntVarP(&port, "port", "", wol.DefaultPort, "Broadcast Port (optional)")
}

func runWakeCmd(cmd *cobra.Command, args []string) {
	packet, err := types.NewMagicPacket(mac)
	if err != nil {
		log.Fatalf("Failed to create Magic packet: %s", err.Error())
	}

	log.Printf("Sending magic packet to MAC %s", mac)
	if err := wol.Send(host, port, packet); err != nil {
		log.Fatalf("Failed to send magic packet: %s", err.Error())
	}
}
