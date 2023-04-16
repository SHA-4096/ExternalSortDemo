package main

import (
	"ExternalSortDemo/pipeline"
	"bufio"
	"fmt"
	"os"
)

func extSort() {
	fileOut, err := os.Open(fileNameIn)
	if err != nil {
		panic(err)
	}
	p := pipeline.CreatePipeline(fileNameIn, length*64, 8)
	writer := bufio.NewWriter(fileOut)
	pipeline.WriteToSink(writer, p)
	fmt.Println("AAA")
	writer.Flush()

}
