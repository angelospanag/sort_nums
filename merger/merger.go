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

// MergeRuns performs a k-way merge of the all the temporary files that hold our sorted runs
func MergeRuns(chunksNum int) error {
	log.Printf("Trying to merge runs of %d chunks\n", chunksNum)

	// Create a priority queue
	pq := make(datastructs.PriorityQueue, chunksNum)

	// Create the output file that will hold the sorted integers
	outputFile, err := fileoperations.CreateFile("sorted_output.txt")
	if err != nil {
		return err
	}
	outputFileWriter := bufio.NewWriter(outputFile)

	// Open a stream for every temporary file we created that holds our sorted runs
	for fileCounter := 0; fileCounter < chunksNum; fileCounter++ {

		// Assuming that our temporary files are of the form tmp_*
		fileName := "tmp_" + strconv.Itoa(fileCounter)
		f, err := fileoperations.OpenFile(&fileName)
		if err != nil {
			return err
		}

		r := bufio.NewReader(f)

		nextInt, err := fileoperations.ReadNextNumFromCSVFile(r)
		if err != nil {
			return err
		}

		// Fill the priority queue with the first set of integers from our
		// temporary files
		queueItem := make(map[*bufio.Reader]int)
		queueItem[r] = nextInt

		pq[fileCounter] = &datastructs.Item{
			Value:    queueItem,
			Priority: nextInt,
			Index:    fileCounter,
		}
	}

	// Initialise the priority queue
	heap.Init(&pq)

	// Main looping logic
	// As long as our queue is not empty, keep taking out the smallest elements
	// in it and write it to the final output file
	for pq.Len() > 0 {
		output := heap.Pop(&pq).(*datastructs.Item)

		// Write the smallest element to the output file
		outputFileWriter.WriteString(strconv.Itoa(output.Priority))
		outputFileWriter.WriteString(",")

		// Next element from the file...
		for reader := range output.Value {

			i, err := fileoperations.ReadNextNumFromCSVFile(reader)
			// End of file?
			if err == io.EOF {
				log.Println(err)
				log.Println(i)
			} else { // ...if there is still more, just push it in the queue
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
