package main

import (
	"ExternalSortDemo/pipeline"
	"bufio"
	"math/rand"
	"os"
)

func RandGen(num int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < num; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

//global vars
func generate() {
	file, err := os.Create(fileNameIn)
	if err != nil {
		panic((err))
	}
	defer file.Close()   //main执行完毕之后才会关闭
	p := RandGen(length) //现在p是一个channel
	writer := bufio.NewWriter(file)
	pipeline.WriteToSink(writer, p)
	writer.Flush()

}
