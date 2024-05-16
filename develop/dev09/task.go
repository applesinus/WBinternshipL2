package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	Ex9("https://ru.m.wikipedia.org/wiki/Wildberries")
}

func Ex9(url string) {
	outputDir := "downloads"

	if url == "" {
		fmt.Println("Please provide a URL")
		return
	}

	downloadWebsite(url, outputDir)
}

func downloadWebsite(url, outputDir string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch URL: %s\n", resp.Status)
		return
	}

	baseURL := strings.TrimSuffix(url, filepath.Ext(url))
	basePath := filepath.Join(outputDir, filepath.Base(baseURL))
	err = os.MkdirAll(basePath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	indexFile, err := os.Create(filepath.Join(basePath, "index.html"))
	if err != nil {
		fmt.Printf("Error creating index file: %v\n", err)
		return
	}
	defer indexFile.Close()

	_, err = io.Copy(indexFile, resp.Body)
	if err != nil {
		fmt.Printf("Error copying response to file: %v\n", err)
		return
	}

	fmt.Printf("Website downloaded successfully to %s\n", basePath)
}
