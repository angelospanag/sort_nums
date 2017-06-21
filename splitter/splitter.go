package splitter

import (
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/angelospanag/sort_nums/converters"
	"github.com/angelospanag/sort_nums/fileoperations"
)

func SplitFileToChunks(file *os.File, bufferSize int64) (int, error) {

	bufferData := make([]byte, bufferSize)
	counter := 0

	for {

		n, err := file.Read(bufferData)
		if err != nil && err != io.EOF {
			return -1, err
		}
		if n == 0 {
			break
		}

		//Index of the last byte of the saved buffer
		index := n - 1

		// We haven't hit a comma... Probably we sliced a number in half
		if string(bufferData[index]) != "," {
			var goBack int64
			// We want to keep the maximum limit of buffer so we calculate how many bytes we should move back
			for string(bufferData[index]) != "," {
				index--
				goBack++
			}
			// Also roll back our file pointer so the next run will start parsing the file from a proper position
			file.Seek(-goBack, 1)
		}

		intSlice, err := converters.CSVStringToIntSlice(string(bufferData[:index]))
		if err != nil {
			return -1, err
		}

		sortedBuffer, err := sortBuffer(intSlice)
		if err != nil {
			return -1, err
		}

		// Create a temporary file and write it
		tmpFilename := "tmp_" + strconv.Itoa(counter)
		counter++

		tmpFile, err := fileoperations.CreateFile(tmpFilename)
		if err != nil {
			return -1, err
		}

		defer tmpFile.Close()
		writeBufErr := fileoperations.WriteBufferToFile(sortedBuffer, tmpFile)
		if writeBufErr != nil {
			return -1, writeBufErr
		}
	}
	return counter, nil
}

func sortBuffer(ints []int) ([]int, error) {
	sort.Ints(ints)
	return ints, nil
}
