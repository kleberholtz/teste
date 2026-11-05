package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

func main() {
	progress := flag.Bool("progress", false, "conta de 0 a 10")
	random := flag.Bool("random", false, "imprime uma string aleatória")
	flag.Parse()

	if *progress {
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
		return
	}

	if *random {
		const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		const length = 16
		result := make([]byte, length)

		for i := range result {
			value, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
			if err != nil {
				fmt.Fprintf(os.Stderr, "erro ao gerar string aleatória: %v\n", err)
				os.Exit(1)
			}
			result[i] = characters[value.Int64()]
		}

		fmt.Println(string(result))
		return
	}

	fmt.Println("hello world")
}
