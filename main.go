package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	progress := flag.Bool("progress", false, "conta de 0 a 10")
	ping := flag.Bool("ping", false, "faz ping para 1.1.1.1")
	flag.Parse()

	if *progress {
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
		return
	}

	if *ping {
		command := exec.Command("ping", "-c", "4", "1.1.1.1")
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "erro ao executar ping: %v\n", err)
			os.Exit(1)
		}
		return
	}

	fmt.Println("hello world")
}
