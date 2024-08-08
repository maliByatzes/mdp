package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "file", "", "Markdown file to preview")
	flag.Parse()

	if fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(fileContent))
}
