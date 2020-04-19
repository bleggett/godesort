package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godesort",
	Short: "Sorts disc images for GDEMU's Rhea and Phoebe Sega Saturn optical disc emulators (ODEs)",
	Long: ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("imageroot", "i", "/media/sdcard", "Path to folder root that contains numbered GDEMU image folders")
}
