package pipeline

import (
	"fmt"
	"sort"
)

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		sort.Ints(a[:])
		fmt.Println("InMemSortDone")
		fmt.Println(a[:])
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out

}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v2
			} else {
				out <- v1
			}
		}
		close(out)
		fmt.Println("Merged a segement")
	}()
	return out
}

func Merge_main(inputs ...<-chan int) <-chan int {
	fmt.Println(inputs)
	if len(inputs) == 1 {
		return inputs[0]
	}
	mid := len(inputs) / 2
	return Merge(
		Merge_main(inputs[:mid]...),
		Merge_main(inputs[mid:]...)) //这里括号还不能换行

}
