package util

import (
	"bufio"
	"path/filepath"
	"log"
	"os"
)

func readOneLineTextFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scanner.Text()
}

func ReadOneLineFileIfExists(folder string, fileName string) string {
	file := filepath.Join(folder, fileName)
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		return readOneLineTextFile(file)
	}
	return ""
}
