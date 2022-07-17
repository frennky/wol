package cmd

import (
	"log"

	"github.com/frennky/wol/pkg/types"
	"github.com/frennky/wol/pkg/wol"
	"github.com/spf13/cobra"
)

var mac string
var ip string
var port int

var wakeCmd = &cobra.Command{
	Use:   "wake",
	Short: "Wake On LAN",
	Long:  `Sends magic packet to target MAC.`,
	Run:   runWakeCmd,
}

func init() {
	rootCmd.AddCommand(wakeCmd)
	wakeCmd.PersistentFlags().StringVarP(&mac, "mac", "m", "", "Target MAC Address (required)")
	wakeCmd.MarkFlagRequired("mac")
	wakeCmd.PersistentFlags().StringVarP(&ip, "ip", "i", wol.DefaultBroadcastIPAddress, "Broadcast IP Address (optional)")
	wakeCmd.PersistentFlags().IntVarP(&port, "port", "p", wol.DefaultPort, "Broadcast Port (optional)")
}

func runWakeCmd(cmd *cobra.Command, args []string) {
	packet, err := types.NewMagicPacket(mac)
	if err != nil {
		log.Fatalf("Failed to create Magic packet: %s", err.Error())
	}

	log.Printf("Sending magic packet to MAC %s", mac)
	if err := wol.Send(ip, port, packet); err != nil {
		log.Fatalf("Failed to send magic packet: %s", err.Error())
	}
}
