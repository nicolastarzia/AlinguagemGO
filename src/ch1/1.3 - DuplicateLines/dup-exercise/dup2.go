package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		totNumber := 0
		strFiles := ""
		strSepFiles := " "
		for line2, n2 := range n {
			totNumber += n2
			strFiles = line2 + strSepFiles
			strSepFiles = ", "
		}
		//if totNumber > 1 {
		fmt.Printf("%d\t%s\t%s\n", totNumber, line, strFiles)
		//}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()][f.Name()]++
	}
}
