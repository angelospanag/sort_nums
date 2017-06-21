package fileoperations

import (
	"bufio"
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
func ReadNextNumFromCSVFile(r *bufio.Reader) (int, error) {

	numberToken := []byte{}
	var numberTokenAsInt int

	for {
		fileByte := make([]byte, 1)

		_, err := r.Read(fileByte)

		// if err == io.EOF {
		// 	return numberTokenAsInt, nil
		// }

		if err != nil {
			return 0, err
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

		// if string(token[0]) != "," {
		// 	numberToken = append(numberToken, token[0])
		// } else {
		// 	numberTokenAsString := string(numberToken)
		// 	numberTokenAsInt, err = strconv.Atoi(numberTokenAsString)
		// 	if err != nil {
		// 		return 0, err
		// 	}
		// 	break
		// }
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
