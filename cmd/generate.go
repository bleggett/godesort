package cmd

import (
	"fmt"
	"log"
	"strings"
	"path/filepath"
	"os"
	"github.com/spf13/cobra"
	"github.com/bleggett/godesort/saturn"
	"github.com/bleggett/godesort/rmenu"
	"github.com/bleggett/godesort/util"
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
		images := scanOrderedRoot(path)
		rmenu.WriteAllDiscInfo(images, path)
		rmenu.RunMkisofs(path)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func scanOrderedRoot(rootPath string) []saturn.SaturnImage {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		log.Fatalf("Path %s does not exist!", rootPath)
	}
	file, err := os.Open(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	files, err := file.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	return analyzeFolders(files, rootPath)
}

func analyzeFolders(subfolders []os.FileInfo, rootPath string) []saturn.SaturnImage {
	images := make([]saturn.SaturnImage, 0)
	for _, file := range subfolders {
		if file.IsDir() {
			globber := filepath.Join(rootPath, file.Name(), "*.img")
			matches, _ := filepath.Glob(globber)
			if len(matches) > 0 {
				res := saturn.ReadDisc_CCD(filepath.Join(matches[0]))
				fmt.Printf("Disc image: %+v \n", res)
				images = append(images, res)
				//Skip the 01/001/etc dir - that's where RMENU lives and we take care of that elsewhere
			} else if strings.HasSuffix(file.Name(), "01") {
				//TODO write enry
				res := buildRMENUEntry()
				images = append(images, res)
			} else {
				//TODO constify title txt
				titleFile := filepath.Join(rootPath, file.Name())
				sep := buildSeparatorEntry(titleFile, "title.txt")
				if sep.Title != "" {
					fmt.Printf("Title file: %s", sep.Title)
					images = append(images, sep)
				}
			}
		}
	}

	return images
}

func buildSeparatorEntry(titlePath string, titlefile string) saturn.SaturnImage {
	title := util.ReadOneLineFileIfExists(titlePath, titlefile)
	imageCountDir := filepath.Base(titlePath)

	return saturn.SaturnImage{
		Title: title,
		DiscNumber: 1,
		DiscCount: 1,
		Order: imageCountDir,
	}
}

func buildRMENUEntry() saturn.SaturnImage {
	return saturn.SaturnImage{
		Title: "RMENU",
		DiscNumber: 1,
		DiscCount: 1,
		Region: "JTUE",
		Version: "V0.2.0",
		Date: "20170228",
		Order: "01",
	}
}
