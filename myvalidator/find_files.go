package main

import (
	"log"
	"os"
	"path/filepath"
)

func findFiles(path string) error {
	if err := filepath.WalkDir(path, walk); err != nil {
		return err
	}

	return nil
}

func walk(path string, d os.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.IsDir() {
		return nil
	}

	if filepath.Ext(path) != ".go" {
		return nil
	}

	structs, err := parseFile(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, structInfo := range structs {
		structInfo.PrintInfo()
	}

	if err := generateCode(structs); err != nil {
		log.Fatal(err)
	}

	return nil
}
