package main

import (
	"log"
	"os"
)

func main() {
	argsWithoutCmd := os.Args[1:]
	if len(argsWithoutCmd) == 0 {
		log.Fatal("Invalid parameters:\n\tvalidatorgen <path>\n")
	}

	if err := findFiles(argsWithoutCmd[0]); err != nil {
		log.Fatal(err)
	}
}
