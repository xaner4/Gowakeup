package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/xaner4/gowakeup/pkg/wol"
)

var (
	port int
	ip   string
)

// Execute the Cobra CLI
func Execute() {

	var cmdWake = &cobra.Command{
		Use:   "wake [MAC/Alias address to device]",
		Short: "Sends a magick packet to the MAC address provided",
		Long: `print is for printing anything back to the screen.
	For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			mp, err := wol.CreateMagicPacket(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			err = wol.SendMagicPacket(mp, ip, port)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Magic packet sent successfully to %q on port %d\n", ip, port)
			}
		},
	}

	cmdWake.PersistentFlags().IntVarP(&port, "port", "p", 9, "Destination port")
	cmdWake.PersistentFlags().StringVarP(&ip, "ip", "i", "255.255.255.255", "Destination IP address")
	var rootCmd = &cobra.Command{Use: "gowakeup"}
	rootCmd.AddCommand(cmdWake)
	rootCmd.Execute()
}
