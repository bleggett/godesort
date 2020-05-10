package rmenu

import (
	"log"
	"os/exec"
	"path/filepath"
)

//I'd like to use a Go library for this, like
// "github.com/KarpelesLab/iso9660" - but the only boot images these
// libs support are El Torito - which is not what we want for the Saturn

func RunMkisofs(rootDir string) {
	mkisofs := "mkisofs"

	rmenuDir := filepath.Join(rootDir, "01", "BIN", "RMENU")
	outputIso := filepath.Join(rootDir, "01", "RMENU.iso")
	cmd := exec.Command(mkisofs, "-quiet", "-sysid", "SEGA SATURN", "-V", "RMENU", "-volset", "RMENU", "-publisher",
		"SEGA ENTERPRISES, LTD.", "-p", "SEGA ENTERPRISES, LTD", "-A", "RMENU", "-G", "IP.BIN", "-l", "-input-charset", "iso8859-1", "-o", outputIso, rmenuDir)

	cmd.Dir = rmenuDir
	_, err := cmd.Output()

	if err != nil {
		log.Println(err.Error())
		return
	}

	// log.Println(string(stdout))
}
