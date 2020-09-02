package main

import (
	"time"

	"github.com/verabull/pb"
)

func main() {
	bar, _ := pb.DefaultLsdBar(
		100,
		pb.SetWidth(100),
		pb.SetTitle("Downloading"),
		pb.SetPrefix('('),
		pb.SetSuffix(')'),
		pb.SetFilledSymbol('='),
		pb.SetUnfilledSymbol(' '),
	)

	for i := 0; i < bar.MaxIteration; i++ {
		bar.Add()
		time.Sleep(30 * time.Millisecond)
	}
}
