package splitter

import (
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/angelospanag/sort_nums/converters"
	"github.com/angelospanag/sort_nums/fileoperations"
)

func SplitFileToChunks(file *os.File, bufferSize int64) error {

	bufferData := make([]byte, bufferSize)
	counter := 0

	for {

		n, err := file.Read(bufferData)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		intSlice, err := converters.CSVStringToIntSlice(string(bufferData[:n]))
		if err != nil {
			return err
		}

		sortedBuffer, err := sortBuffer(intSlice)
		if err != nil {
			return err
		}

		// Create a temporary file and write it

		tmpFilename := "tmp_" + strconv.Itoa(counter)
		counter++

		tmpFile, err := fileoperations.CreateFile(tmpFilename)
		if err != nil {
			return err
		}

		defer tmpFile.Close()
		writeBufErr := fileoperations.WriteBufferToFile(sortedBuffer, tmpFile)
		if writeBufErr != nil {
			return writeBufErr
		}
	}
	return nil
}

func sortBuffer(ints []int) ([]int, error) {
	sort.Ints(ints)
	return ints, nil
}
