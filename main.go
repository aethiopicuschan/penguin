package main

import (
	"fmt"
	"os"

	"github.com/aethiopicuschan/penguin/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	fmt.Println("Done!")
}
