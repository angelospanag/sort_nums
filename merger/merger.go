package merger

import (
	"bufio"
	"container/heap"
	"io"
	"log"
	"strconv"

	"github.com/angelospanag/sort_nums/datastructs"
	"github.com/angelospanag/sort_nums/fileoperations"
)

// MergeRuns
func MergeRuns(chunksNum int) error {
	log.Printf("Trying to merge runs of %d chunks\n", chunksNum)
	pq := make(datastructs.PriorityQueue, chunksNum)

	outputFile, err := fileoperations.CreateFile("sorted_output.txt")
	if err != nil {
		return err
	}
	outputFileWriter := bufio.NewWriter(outputFile)

	for fileCounter := 0; fileCounter < chunksNum; fileCounter++ {

		fileName := "tmp_" + strconv.Itoa(fileCounter)
		f, err := fileoperations.OpenFile(&fileName)
		if err != nil {
			return err
		}

		r := bufio.NewReader(f)

		//fileToReaderMapper[fileCounter] = r
		nextInt, err := fileoperations.ReadNextNumFromCSVFile(r)

		queueItem := make(map[*bufio.Reader]int)
		queueItem[r] = nextInt

		pq[fileCounter] = &datastructs.Item{
			Value:    queueItem,
			Priority: nextInt,
			Index:    fileCounter,
		}
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		output := heap.Pop(&pq).(*datastructs.Item)

		// Write the sorted element to our output file
		outputFileWriter.WriteString(strconv.Itoa(output.Priority))
		outputFileWriter.WriteString(",")

		for reader := range output.Value {

			i, err := fileoperations.ReadNextNumFromCSVFile(reader)

			if err == io.EOF {
				log.Println(err)
				log.Println(i)
			} else {
				queueItem := make(map[*bufio.Reader]int)
				queueItem[reader] = i
				item := &datastructs.Item{Value: queueItem, Priority: i}
				heap.Push(&pq, item)
			}
		}
	}
	outputFileWriter.Flush()

	return nil
}
