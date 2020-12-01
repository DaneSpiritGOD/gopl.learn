// 练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
package main

import (
	"fmt"
	"io"
	"os"
)

type countingWriter struct {
	writeN int64
	wr     io.Writer
}

// CountingWriter return a wrapped io.Writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := countingWriter{
		wr: w,
	}
	return &newWriter, &newWriter.writeN
}

func (w *countingWriter) Write(p []byte) (nn int, err error) {
	nn, err = w.wr.Write(p)
	w.writeN += int64(nn)
	return
}

func main() {
	writer, pn := CountingWriter(os.Stdout)
	fmt.Printf("new count: %d\n", *pn)

	fmt.Fprintf(writer, "hello, %s\n", "jack")
	fmt.Printf("after write string, new count: %d", *pn)
}
