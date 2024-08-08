package main

import (
	"bytes"
	"os"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func TestRun(t *testing.T) {
	if err := run(inputFile, true); err != nil {
		t.Error(err)
	}

	resultContent, err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}

	expectedContent, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}

	if !bytes.Equal(resultContent, expectedContent) {
		t.Logf("resultContent: %s", resultContent)
		t.Logf("expectedContent: %s", expectedContent)
		t.Error("Result file content does not match expected file content")
	}
}

func TestConvertToHTML(t *testing.T) {
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}

	resultContent := convertToHTML(fileContent)

	expectedContent, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}

	if !bytes.Equal(resultContent, expectedContent) {
		t.Logf("resultContent: %s", resultContent)
		t.Logf("expectedContent: %s", expectedContent)
		t.Error("Result file content does not match expected file content")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(inputFile, true)
	}
}

func BenchmarkConvertToHTML(b *testing.B) {
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		b.Errorf("Error reading file: %s", err)
	}

	for i := 0; i < b.N; i++ {
		convertToHTML(fileContent)
	}
}
