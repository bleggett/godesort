package rmenu

import (
	"fmt"
	"log"
	"path/filepath"
	"os"
	"github.com/bleggett/godesort/saturn"
)

var inifile string = "BIN/RMENU/LIST.INI"

func WriteAllDiscInfo(images []saturn.SaturnImage, rootPath string) {
	ini := filepath.Join(rootPath, "01", inifile)
	if _, err := os.Stat(filepath.Dir(ini)); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(ini), 0700)
	}
	file, err := os.Create(ini)
	checkFileErr(err)
	defer file.Close()

	for _, image := range images {
		writeDiscInfo(file, image)
	}

	log.Printf("Wrote data for %d images to %s", len(images), ini)

}

func writeDiscInfo(file *os.File, image saturn.SaturnImage) {
	writeLine(file, image.Order, "title", image.Title)
	writeLine(file, image.Order, "disc", fmt.Sprintf("%d/%d", image.DiscNumber, image.DiscCount))
	writeLine(file, image.Order, "region", image.Region)
	writeLine(file, image.Order, "version", image.Version)
	writeLine(file, image.Order, "date", image.Date)
}

func writeLine(file *os.File, numberDir string, key string, value string) {
	_, err := file.WriteString(fmt.Sprintf("%s.%s=%s\n", numberDir, key, value))
	checkFileErr(err)
}

func checkFileErr(e error) {
    if e != nil {
        panic(e)
    }
}
