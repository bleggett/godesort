package rmenu

import (
	"fmt"
	"os"
	"path/filepath"
	"log"
	"bufio"
)

var imageExt string = "*.ccd"

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
		tag := getTagIfExist(filepath.Dir(match))

		log.Printf("Found image: %s", match)
		imgGroups[tag] = append(imgGroups[tag], ImageSet{filepath.Dir(match), filepath.Base(match), tag})
	}

	return imgGroups
}

func getTagIfExist(folder string) string {
	tagFile := filepath.Join(folder, "tag.txt")
	if _, err := os.Stat(tagFile); !os.IsNotExist(err) {
		return readTag(tagFile)
	}
	return ""
}

func readTag(tagFile string) string {
	file, err := os.Open(tagFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	tagName := scanner.Text()
	return tagName
}

//INIScanner
//TODO
//For every numbered subfolder in the root
//Either look for an .img file
//or look for a TITLE.TXT file, or whatever
