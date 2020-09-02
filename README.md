# Installation  
```sh
$ go get github.com/verabull/pb     
```  

# Quick start   
```golang
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
```  

# Сustomization  
```golang
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
```  

# Attention  
I do not promote drug use

# LICENCE  
MPL-2.0 License  
