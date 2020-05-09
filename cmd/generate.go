package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"os"
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
		path, _ := cmd.Flags().GetString("imageroot")
		fmt.Printf("generate called on %s\n", path)
		images := readAllDiscInfo(path)
		rmenu.WriteAllDiscInfo(images)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func readAllDiscInfo(path string) []saturn.SaturnImage {
	images := make([]saturn.SaturnImage, 1)

	imgSetGroups := rmenu.BuildMap(path)

	for _, imgSet := range imgSetGroups {
		for _, img := range imgSet {
			res := saturn.ReadDisc_CCD(filepath.Join(img.SourceDir, "/", img.ImageName))
			fmt.Printf("Disc image: %+v \n", res)
			images = append(images, res)
		}
	}

	return images
}

func scanOrderedRoot(rootPath string) {
	file, err := os.Open(rootPath)
	if err != nil {
		log.Fatal(err)
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(names)
}
//INIScanner
//TODO
//For every numbered subfolder in the root
//Either look for an .img file
//or look for a TITLE.TXT file, or whatever
