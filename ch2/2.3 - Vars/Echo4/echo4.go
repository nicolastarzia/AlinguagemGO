package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "nao envia a quebra de linha")
var sep = flag.String("s", " ", "separador")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
