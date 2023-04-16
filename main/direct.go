package main

import "fmt"

const fileNameOut = "sorted.out"
const fileNameIn = "data.in"
const length = 64

func main() {
	for {
		op := 0
		fmt.Println("1:generate 2:readSorted 3:sort 4:quit")
		fmt.Scanf("%d", op)
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
