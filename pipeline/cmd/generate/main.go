package main

import (
	"fmt"
	"os"

	"github.com/digiconvent/seeds/internal/tree"
)

func main() {
	os.Chdir("..")
	err := tree.Load("seeds").Validate().Generate().Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("index.json and readme regenerated")
}
