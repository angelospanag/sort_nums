package main

import (
	"flag"
	"log"
	"math"

	"github.com/angelospanag/sort_nums/fileoperations"
	"github.com/angelospanag/sort_nums/merger"
	"github.com/angelospanag/sort_nums/splitter"
)

func main() {

	var err error
	filePtr := flag.String("file", "", "Input file")
	memoryPtr := flag.Int64("memory", 0, "Memory (in bytes) that will be used to hold and sort a series of numbers")

	flag.Parse()

	if *memoryPtr <= 0 {
		log.Fatal("Not valid memory value")
	}

	inputFile, err := fileoperations.OpenFile(filePtr)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := inputFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var chunksAsFloat float64 = float64(fi.Size()) / float64(*memoryPtr)
	chunksNum := int64(math.Ceil(chunksAsFloat))

	log.Printf("File %s is %d bytes, will be split to %d chunks\n", *filePtr, fi.Size(), chunksNum)

	// Split file to chunks and store them in disk
	err = splitter.SplitFileToChunks(inputFile, *memoryPtr)
	if err != nil {
		log.Fatal(err)
	}

	// Merge chunks stored in memory in one output file
	err = merger.MergeRuns(int(chunksNum))
	if err != nil {
		log.Fatal(err)
	}

	// Remove temporary files from previous runs
	err = fileoperations.CleanupTempFiles(int(chunksNum))
	if err != nil {
		log.Fatal(err)
	}
}
