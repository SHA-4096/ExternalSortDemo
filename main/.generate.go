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
const FileName = "data.in"
const length = 64

func main() {
	file, err := os.Create(FileName)
	if err != nil {
		panic((err))
	}
	defer file.Close()   //main执行完毕之后才会关闭
	p := RandGen(length) //现在p是一个channel
	writer := bufio.NewWriter(file)
	pipeline.WriteToFile(writer, p)
	writer.Flush()
	file, err = os.Open(FileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

}
