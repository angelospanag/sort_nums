package fileoperations

import (
	"io"
	"os"
	"strconv"

	"github.com/angelospanag/sort_nums/converters"
)

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

		if string(fileByte[0]) != "," {
			numberToken = append(numberToken, fileByte[0])
		} else {
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
