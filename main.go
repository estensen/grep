package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("usage: grep <pattern> <filename>")
	}

	pattern := args[0]
	fileName := args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("cannot read file", err)
	}
	//defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			log.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

