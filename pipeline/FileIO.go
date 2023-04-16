package pipeline

import (
	"encoding/binary"
	"io"
)

func WriteToFile(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8) //byte是无符号8位整型，实际上一次写了64位
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func ReadFromFile(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read((buffer))
			bytesRead += n
			if n > 0 {
				out <- int(binary.BigEndian.Uint64((buffer)))
			}
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out) //不写这个会报错deadlock
	}()
	return out
}
