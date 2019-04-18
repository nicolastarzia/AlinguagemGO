package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t0 := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	t1 := time.Now()
	fmt.Println(s)
	fmt.Printf("The call took %.20fs to run.\n", t1.Sub(t0).Seconds())
}
