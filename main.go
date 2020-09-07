package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("missing filename arg")
	}

	fileName := args[0]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("cannot read file", err)
	}
	defer file.Close()
}
