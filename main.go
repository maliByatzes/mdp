package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

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

	outName := filepath.Base(fileName) + ".html"
	// fmt.Println(outName)

	return saveToHTMLFile(outName, htmlContent)
}

func convertToHTML(content []byte) []byte {
	output := blackfriday.Run(content)
	htmlContent := bluemonday.UGCPolicy().SanitizeBytes(output)

	return htmlContent
}

func saveToHTMLFile(outName string, content []byte) error {
	return os.WriteFile(outName, content, 0644)
}
