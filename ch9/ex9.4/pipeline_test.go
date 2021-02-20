package main

import (
	"testing"
)

func BenchmarkPipeline(b *testing.B) {
	in, out := makePipeline(10000000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}
