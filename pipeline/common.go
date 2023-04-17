package pipeline

import (
	"bufio"
	"fmt"
	"os"
)

func CreatePipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	sortRes := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := ReadFromFile(bufio.NewReader(file), chunkSize)
		sortRes = append(sortRes, InMemSort(source))
		fmt.Println(sortRes[:])

	}
	return Merge_main(sortRes...)
}
