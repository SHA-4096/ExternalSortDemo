package main

import (
	"ExternalSortDemo/pipeline"
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const fileNameOut = "sorted.out"
const fileNameIn = "data.in"
const length = 64

func main() {
	for {
		op := 0
		fmt.Println("1:generate 2:readSorted 3:sort 4:quit")
		fmt.Scanf("%d", &op)
		switch op {
		case 1:
			generate()
		case 2:
			readFile()
		case 3:
			extSort()
		default:
			return
		}

	}
}

func extSort() {
	fileOut, err := os.Create(fileNameOut) //不能用Open啊……Open默认是只读的啊……
	if err != nil {
		panic(err)
	}
	p := pipeline.CreatePipeline(fileNameIn, length*64, 4)
	writer := bufio.NewWriter(fileOut)
	pipeline.WriteToSink(writer, p)
	fmt.Println("AAA")
	writer.Flush()

}

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
