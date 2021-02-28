package main

import (
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mtype path_to_file")
		os.Exit(1)
	}
	m, err := mimetype.DetectFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(m)
}
