package fileoperations

import (
	"io"
	"os"
	"strconv"

	"github.com/angelospanag/sort_nums/converters"
)

//CreateFile creates a file given a path in the local filesystem
func CreateFile(filename string) (*os.File, error) {
	f, err := os.Create(filename)

	if err != nil {
		return nil, err
	}
	return f, nil
}

// OpenFile opens a file
func OpenFile(filePath *string) (*os.File, error) {
	inputFile, err := os.Open(*filePath)

	if err != nil {
		return nil, err
	}
	return inputFile, nil
}

// WriteBufferToFile writes a slice of integers to a specified file in CSV format
func WriteBufferToFile(ints []int, file *os.File) error {

	b, err := converters.IntSliceToCSVString(ints)
	if err != nil {
		return err
	}
	file.WriteString(b)

	return nil
}

// ReadNextNumFromCSVFile parses a CSV file using its already open reader and returns the next integer
func ReadNextNumFromCSVFile(f *os.File) (int, error) {

	numberToken := []byte{}
	var numberTokenAsInt int

	for {
		fileByte := make([]byte, 1)

		_, err := f.Read(fileByte)

		// IMPORTANT! Corner case: when we have reached EOF from a previous run
		// we still have leftover buffer to return.
		// Check if that's the case and return it
		if err == io.EOF && len(numberToken) > 0 {
			lastNumberTokenAsString := string(numberToken)
			lastNumberTokenAsInt, errConv := strconv.Atoi(lastNumberTokenAsString)
			if errConv != nil {
				return -1, err
			}
			return lastNumberTokenAsInt, nil

		} else if err != nil {
			return -1, err
		}

		// If a read character is NOT a comma, we still have numbers to append
		if string(fileByte[0]) != "," {
			numberToken = append(numberToken, fileByte[0])
		} else { //...we found a comma, we have a complete number to return
			numberTokenAsString := string(numberToken)
			numberTokenAsInt, err = strconv.Atoi(numberTokenAsString)
			if err != nil {
				return 0, err
			}
			break
		}
	}

	return numberTokenAsInt, nil
}

// CleanupTempFiles removes any temporary files created from previous runs
func CleanupTempFiles(fileNum int) error {

	for counter := 0; counter < fileNum; counter++ {
		err := os.Remove("tmp_" + strconv.Itoa(counter))
		if err != nil {
			return err
		}
	}

	return nil
}
