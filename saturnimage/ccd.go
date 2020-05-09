package saturnimage

import (
	"bufio"
	"io"
	"os"
	"log"
	"strings"
	"strconv"
)

func ReadDisc_CCD(imgName string) SaturnImage {
	fd, err := os.Open(imgName)
	if err != nil { //error handler
		log.Fatalf("Error opening image file to parse metadata, error was: %s", err)
	}
	defer fd.Close()

    number, count := getDiscNumber_CCD(fd)

	return SaturnImage{
		title: getDiscTitle_CCD(fd),
		discNumber: number,
		discCount: count,
		discRegion: getDiscRegion_CCD(fd),
		discVersion: getDiscVersion_CCD(fd),
		discDate: getDiscDate_CCD(fd),
	}
}

func getDiscTitle_CCD(fd *os.File) string {
	title := getStringAtOffset(fd, 112, 55)
	return title
}

func getDiscNumber_CCD(fd *os.File) (int, int) {
	counts := strings.Split(getStringAtOffset(fd, 75, 3), "/")
	number, _ := strconv.Atoi(counts[0])
	count, _ := strconv.Atoi(counts[1])
	return number, count
}

func getDiscRegion_CCD(fd *os.File) string {
	region := getStringAtOffset(fd, 80, 10)
	return region
}

func getDiscVersion_CCD(fd *os.File) string {
	version := getStringAtOffset(fd, 59, 5)
	return version
}

func getDiscDate_CCD(fd *os.File) string {
	date := getStringAtOffset(fd, 64, 8)
	return date
}

func getStringAtOffset(fd *os.File, offset int64, size int) string {
	byteCount := make([]byte, size)
	fd.Seek(offset, 0)
	reader := bufio.NewReader(fd) // creates a new reader
	n, _ := io.ReadFull(reader, byteCount)
	if n < size {
		log.Fatalf("Error reading image file at offset %d - expected %d bytes but got %d", offset, size, n)
	}

	return strings.TrimSpace(string(byteCount[:]))
}
