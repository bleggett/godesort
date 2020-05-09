package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/bleggett/godesort/saturn"
	"github.com/bleggett/godesort/rmenu"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		res := saturn.ReadDisc_CCD("/Volumes/Games/test/05/Battle Garegga (Japan).img")
		discs := make([]saturn.SaturnImage, 1)
		discs[0] = res
		fmt.Printf("Disc image: %+v \n", res)
		rmenu.WriteAllDiscInfo(discs)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
