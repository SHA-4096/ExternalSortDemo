package main

import (
	"ExternalSortDemo/pipeline"
	"bufio"
	"os"
)

const FileName = "data.in"

func main() {
	file, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	p := pipeline.ReadFromFile(reader, -1)
	//tmp := []<-chan int{}
	pipeline.Merge(p,)
	//	tmp = append(tmp, p)
	//	out := pipeline.Merge_main(tmp...)
	writer := bufio.NewWriter(file)
	pipeline.WriteToFile(writer, out)
	writer.Flush()

}
