package cmd

import (
	"fmt"

	"gitlab.com/xaner4/gowakeup/pkg/alias"

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
			var mac string
			if e, o := alias.Exists(args[0], args[0]); !e {
				mac = args[0]
			} else {
				ms := alias.Aliases[o]
				mac = ms.Mac
			}

			mp, err := wol.CreateMagicPacket(mac)
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

	var cmdAlias = &cobra.Command{
		Use:   "alias",
		Short: "Lists all Aliases that already exsists",
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			a := alias.Aliases
			if len(a) == 0 {
				fmt.Println("No Aliases exsists yet!\nAdd one with 'alias add [alias] [mac]'")
			}
			for _, v := range a {
				fmt.Printf("Alias: %s is %s mac address", v.Name, v.Mac)
			}
		},
	}

	var cmdAliasAdd = &cobra.Command{
		Use:   "add [name] [mac]",
		Short: "Adds a alias",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := alias.Add(args[0], args[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			alias.Write()
		},
	}

	var cmdAliasRemove = &cobra.Command{
		Use:   "remove [name/mac]",
		Short: "Removes a alias from the aliases",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := alias.Remove(args[0], "")
			if err != nil {
				fmt.Println(err)
				return
			}
			alias.Write()
		},
	}

	cmdWake.PersistentFlags().IntVarP(&port, "port", "p", 9, "Destination port")
	cmdWake.PersistentFlags().StringVarP(&ip, "ip", "i", "255.255.255.255", "Destination IP address")

	cmdAlias.AddCommand(cmdAliasAdd, cmdAliasRemove)

	var rootCmd = &cobra.Command{Use: "gowakeup"}
	rootCmd.AddCommand(cmdWake, cmdAlias)
	rootCmd.Execute()
}
