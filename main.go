package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "file", "", "Markdown file to preview")
	flag.Parse()

	if fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(fileName); err != nil {
		log.Fatal(err)
	}
}

func run(fileName string) error {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	htmlContent := convertToHTML(fileContent)

	fmt.Println(string(htmlContent))

	// save htmlContent to some file

	return nil
}

func convertToHTML(content []byte) []byte {
	output := blackfriday.Run(content)
	htmlContent := bluemonday.UGCPolicy().SanitizeBytes(output)

	return htmlContent
}
