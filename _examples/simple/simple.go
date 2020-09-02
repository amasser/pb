package main

import (
	"time"

	"github.com/verabull/pb"
)

func main() {
	bar, _ := pb.DefaultLsdBar(100)
	for i := 0; i < bar.MaxIteration; i++ {
		bar.Add()
		time.Sleep(30 * time.Millisecond)
	}
}
