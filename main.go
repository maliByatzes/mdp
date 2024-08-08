package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

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

	f, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}

	outName := f.Name()
	fmt.Println(f.Name())

	if err := saveToHTMLFile(outName, htmlContent); err != nil {
		return err
	}

	defer os.Remove(outName)

	return openPreview(outName)
}

func convertToHTML(content []byte) []byte {
	output := blackfriday.Run(content)
	htmlContent := bluemonday.UGCPolicy().SanitizeBytes(output)

	return htmlContent
}

func saveToHTMLFile(outName string, content []byte) error {
	return os.WriteFile(outName, content, 0644)
}

func openPreview(fileName string) error {
	cName := ""
	cParams := []string{}

	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	cParams = append(cParams, fileName)

	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	err = exec.Command(cPath, cParams...).Run()

	time.Sleep(2 * time.Second)

	return err
}
