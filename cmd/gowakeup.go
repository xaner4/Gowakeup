package cmd

import (
	"github.com/spf13/cobra"
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

		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdWake)
	rootCmd.Execute()
}
