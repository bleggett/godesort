package rmenu

import (
	"fmt"
	"log"
	"os"
	"github.com/bleggett/godesort/saturn"
)

//https://github.com/tehnoir/RheaMenu-macOS/blob/master/rmenu
var inifile string = "./RMENU/LIST.INI"

func WriteAllDiscInfo(images []saturn.SaturnImage) {
	if _, err := os.Stat("./RMENU"); os.IsNotExist(err) {
		os.MkdirAll("./RMENU", 0700)
	}
	file, err := os.Create(inifile)
	checkFileErr(err)
	defer file.Close()

	writeRMENUEntry(file)

	for _, image := range images {
		writeDiscInfo(file, image)
	}

	log.Printf("Wrote data for %d images to %s", len(images), inifile)

}

func writeRMENUEntry(file *os.File) {
	writeLine(file, "01", "title", "RMENU")
	writeLine(file, "01", "disc", "1/1")
	writeLine(file, "01", "region", "JTUE")
	writeLine(file, "01", "version", "V0.2.0")
	writeLine(file, "01", "date", "20170228")
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
