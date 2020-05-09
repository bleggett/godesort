package cmd

import (
	"fmt"
	"sort"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"path"
	"path/filepath"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
    "github.com/bleggett/godesort/util"
)


var separatorTextFile string = "title.txt"

var imageExt string = "*.ccd"
var tagFile string = "tag.txt"

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Reorganizes and orders image subfolders",
	Long: `Looks for CD images in subfolders within the provided image root, and orders
them A-Z by image filename.

Also supports grouping by tags - place a 'tags.txt' file with a single line of text
in every image subfolder you wish to group together.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("imageroot")
		fmt.Printf("sort called on %s\n", path)
		imgGrps := BuildMap(path)
		imgGrps = sortImageGroups(imgGrps)
		tmpSuffix := tempRenameSortedImageSets(path, imgGrps)
		finalRenameSortedDirs(path, tmpSuffix)
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)
}

type ImageSet struct {
	SourceDir string
	ImageName string
	GroupTag string
}


func BuildMap(rootPath string) map[string][]ImageSet {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		log.Fatalf("Path %s does not exist!", rootPath)
	}

	imgGroups := make(map[string][]ImageSet, 0)

	globber := filepath.Join(rootPath, "**", imageExt)
	fmt.Println("Globbing on: ", globber)
	matches, _ := filepath.Glob(globber)
	for _, match := range matches {
		tag := util.ReadOneLineFileIfExists(filepath.Dir(match), tagFile)

		log.Printf("Found image: %s", match)
		imgGroups[tag] = append(imgGroups[tag], ImageSet{filepath.Dir(match), filepath.Base(match), tag})
	}

	return imgGroups
}

func sortImageGroups(imgGrps map[string][]ImageSet) map[string][]ImageSet {

	for grpTag, _ := range imgGrps {
		sort.Slice(imgGrps[grpTag], func(first, second int) bool {
			return imgGrps[grpTag][first].ImageName < imgGrps[grpTag][second].ImageName })
	}

	return imgGrps
}

func tempRenameSortedImageSets(rootDir string, imgGrps map[string][]ImageSet) string {
	guid := xid.New()
	tempPostfix := fmt.Sprintf("-%s", guid.String())
	var counter int = 2
	ungroupedImgSet := imgGrps[""]
	delete(imgGrps, "")

	for _, iSet := range ungroupedImgSet {
		newDir := path.Join(rootDir, fmt.Sprintf("%02d%s", counter, tempPostfix))

		counter++

		os.Rename(iSet.SourceDir,newDir)
		fmt.Printf("Renaming %s and related files\n", strings.TrimSuffix(iSet.ImageName, path.Ext(iSet.ImageName)))
	}

	for grpTag, imgSet := range imgGrps {
		grpPath := path.Join(rootDir, fmt.Sprintf("%02d%s", counter, tempPostfix))
		makeGroupDir(grpPath, grpTag)
		fmt.Printf("Creating Tag Group: %s\n", grpTag)
		counter++
		
		for _, iSet := range imgSet {
			newDir := path.Join(rootDir, fmt.Sprintf("%02d%s", counter, tempPostfix))

			counter++

			os.Rename(iSet.SourceDir,newDir)
			fmt.Printf("Renaming %s and adding to the %s group \n", strings.TrimSuffix(iSet.ImageName, path.Ext(iSet.ImageName)), grpTag)
		}
	}

	return tempPostfix
}

func makeGroupDir(grouppath string, grouptag string) {
	//TODO handle error
	os.Mkdir(grouppath, 0755)
	data := []byte(fmt.Sprintf("%s\n", grouptag))
	//TODO handle error
	ioutil.WriteFile(filepath.Join(grouppath, separatorTextFile), data, 0644)
}

func finalRenameSortedDirs(rootDir string, tmpPostfix string) {
	globber := fmt.Sprintf("%s/*%s", rootDir, tmpPostfix)
	matches, _ := filepath.Glob(globber)
	for _, match := range matches {
		trimPath := strings.TrimSuffix(match, tmpPostfix)
		if _, err := os.Stat(trimPath); !os.IsNotExist(err) {
			//TODO handle remove error
			os.RemoveAll(trimPath)
		}
		os.Rename(match, trimPath)
	}
}


//TODO <optionally test if current path is correct and skip>
//TODO <optionally check for "tag.txt" in imgdir and do grouping based on it
// For the group sort, use a mapof rmenu.ImageSet arrays, where the key is the group tag
// then they can be individually sorted and then concatenated
