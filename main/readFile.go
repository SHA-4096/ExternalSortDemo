package main

//这是一个读取文件的demo
import (
	"ExternalSortDemo/pipeline"
	"bufio"
	"fmt"
	"os"
)

func readFile() {
	file, err := os.Open(fileNameOut)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	p := pipeline.ReadFromFile(reader, -1)
	cnt := 0
	for v := range p {
		fmt.Println(v)
		cnt++

	}

}
