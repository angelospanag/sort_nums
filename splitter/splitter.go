package splitter

import (
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/angelospanag/sort_nums/converters"
	"github.com/angelospanag/sort_nums/fileoperations"
)

// SplitFileToChunks splits a 'comma separated values' file of integers into smaller chunks stored on temporary files
func SplitFileToChunks(file *os.File, bufferSize int64) (int, error) {

	bufferData := make([]byte, bufferSize)

	// This counter will keep track of how many temporary files will be
	// created to keep out sorted buffers in 'slow memory' (hard disk)
	counter := 0

	for {
		// Attempt to read to a buffer
		n, err := file.Read(bufferData)
		if err != nil && err != io.EOF {
			return -1, err
		}
		if n == 0 {
			break
		}

		//Index of the last byte of the saved buffer
		index := n - 1

		var goBack int64
		//We haven't hit a comma...There is a chance that we sliced a number in half
		if string(bufferData[index]) != "," {
			// We don't want to overshoot the maximum limit of buffer so we calculate
			// how many bytes we should move back
			for string(bufferData[index]) != "," && index > 0 {
				index--
				goBack++
			}

			// Also roll back our file pointer so the next run will start parsing the file from a proper position
			file.Seek(-goBack, 1)

			// IMPORTANT! If at some point we reach index = 0 that means that there is
			// only a single integer left to read and no commas are left.
			// In that case just output the remaining number to a new file
			if index == 0 {
				intSlice, err := converters.CSVStringToIntSlice(string(bufferData[0:n]))
				if err != nil {
					return -1, err
				}

				// Create a temporary file for this last number and write it
				// No need to sort a single integer
				tmpFilename := "tmp_" + strconv.Itoa(counter)
				counter++

				tmpFile, err := fileoperations.CreateFile(tmpFilename)
				if err != nil {
					return -1, err
				}

				defer tmpFile.Close()
				writeBufErr := fileoperations.WriteBufferToFile(intSlice, tmpFile)
				if writeBufErr != nil {
					return -1, writeBufErr
				}
				break
			}
		}

		// Convert our read buffer to a slice of integers that will be sorted
		intSlice, err := converters.CSVStringToIntSlice(string(bufferData[:int64(n)-goBack]))
		if err != nil {
			return -1, err
		}

		// Sort it
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

// sortBuffer sorts a slice of integers and returns it
func sortBuffer(ints []int) ([]int, error) {
	sort.Ints(ints)
	return ints, nil
}
