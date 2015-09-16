// before run `go run demo.go`
// 1. you need move demo.go out of this directory
// 2. modify bin_path to your own bin path
package main

import (
	wordmodel "./go-word2vec"
	"bufio"
	"fmt"
	"os"
)

const (
	// change it to your own bin path
	bin_path string = "/Volumes/mini/word2vec/vectors.bin"
)

func main() {
	// load model
	fmt.Println("word2vec model is loading ...")
	model, err := wordmodel.Load(bin_path)
	if err != nil {
		panic(err)
	}

	// read input from console
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Query String: ")
		text, _ := reader.ReadString('\n')

		// query
		pairs, err := model.Query(text, 20)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("|%-20s|%-10s|\n", "WORD:", "DISTANCE:")
			for _, each := range pairs {
				fmt.Printf("|%-20s|%-10.3f|\n", each.Word, each.Sim)
			}
		}
	}
}
