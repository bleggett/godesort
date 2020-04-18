package cmd

import (
	"fmt"
	"log"
	"sort"
	"os"
	"strings"
	"path"
	"path/filepath"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
)


type ImageSet struct {
	SourceDir string
	ImageName string
	SortID int
}

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("imageroot")
		fmt.Printf("sort called on %s\n", path)
		imgSet := buildMap(path)
		imgSet = sortImageSet(imgSet)
		tmpSuffix := tempRenameSortedImageSet(path, imgSet)
		finalRenameSortedDirs(path, tmpSuffix)
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func buildMap(rootPath string) []ImageSet {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		log.Fatalf("Path %s does not exist!", rootPath)
	}

	var imgSet []ImageSet

	//TODO Glob here - walk is a waste
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure walking a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".ccd" {

			imgSet = append(imgSet, ImageSet{filepath.Dir(path), filepath.Base(info.Name()), 0})

			fmt.Printf("Adding CCD: %s at path %s\n", info.Name(), filepath.Dir(path))
			return nil
		}
		// fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", rootPath, err)
		return nil
	}

	return imgSet
}


func sortImageSet(imgSet []ImageSet) []ImageSet {
	sort.Slice(imgSet, func(first, second int) bool {
		return imgSet[first].ImageName < imgSet[second].ImageName })

	fmt.Println("By name", imgSet)
	return imgSet
}

func tempRenameSortedImageSet(rootDir string, imgSet []ImageSet) string {
	guid := xid.New()
	tempPostfix := fmt.Sprintf("-%s", guid.String())

	for i, iSet := range imgSet {
		//TODO Pad name iterator with leading zeroes
		newDir := path.Join(rootDir, fmt.Sprintf("%02d%s", (i+2), tempPostfix))

		// if _, err := os.Stat(newDir); !os.IsNotExist(err) {
		//     // err := os.Rename(newDir, fmt.Sprintf("%s-old", newDir))
		// 	fmt.Printf("Clearing out existing %s to %s-old\n", newDir, newDir)
		// }
		fmt.Printf("Renaming %s to %s\n", iSet.SourceDir, newDir)
		os.Rename(iSet.SourceDir,newDir)
	}

	return tempPostfix
}

func finalRenameSortedDirs(rootDir string, tmpPostfix string) {
	globber := fmt.Sprintf("%s/*%s", rootDir, tmpPostfix)
	fmt.Printf("Globbing on %s\n", globber)
	matches, _ := filepath.Glob(globber)
    for _, match := range matches {
		trimPath := strings.TrimSuffix(match, tmpPostfix)
		fmt.Printf("Dropping postfix - moving %s to %s", match, trimPath)
		os.Rename(match, trimPath)
    }
}
//scan dir
//build map of path/imagename

//Sort alphabetically based on imagename
//rename all to correctfolderid-<UNIQUE SLUG>
//<optionally test if current path is correct and skip>
//<optionally check for "tag.txt" in imgdir and do grouping based on it
//go back and strip -<UNIQUESLUG> from all foldernames
