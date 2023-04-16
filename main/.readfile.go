package main

//这是一个读取文件的demo
import (
	"ExternalSortDemo/pipeline"
	"bufio"
	"fmt"
	"os"
)

const FileName = "data.in"
const length = 64

func main() {
	file, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	p := pipeline.ReadFromFile(reader, -1)
	cnt := 0
	for v := range p {
		fmt.Println(v)
		cnt++
		if cnt > length {
			break //不设置长度限制的时候会有fatal error: all goroutines are asleep - deadlock!
			//但是它是怎么判断go to sleep的？时间吗

		}

	}

}
