package main

import (
	"flag"
	"fmt"
)

func main() {
	progress := flag.Bool("progress", false, "conta de 0 a 10")
	flag.Parse()

	if *progress {
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
		return
	}

	fmt.Println("hello world")
}
