package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

func listMarkdownFiles(dirPath string) ([]string, error) {
	var markdownFiles []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".md" {
			relativePath, err := filepath.Rel(dirPath, path)
			if err != nil {
				return err
			}
			markdownFiles = append(markdownFiles, relativePath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return markdownFiles, nil
}
func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	mdFiles, err := listMarkdownFiles("./md")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte("<h1>Markdown Files</h1><ul>")); err != nil {
		log.Println("Error writing to response:", err)
	}
	for _, file := range mdFiles {
		if _, err := w.Write([]byte(fmt.Sprintf("<li><a href=\"/md//%s\">%s</a></li>", file, file))); err != nil {
			log.Println("Error writing to response:", err)
		}
	}
	if _, err := w.Write([]byte("</ul>")); err != nil {
		log.Println("Error writing to response:", err)
	}
}

func readMarkdownFile(filePath string) ([]byte, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}

func convertMarkdownToHTML(markdownBytes []byte) []byte {
	return markdown.ToHTML(markdownBytes, nil, nil)
}

func serveMarkdownFile(w http.ResponseWriter, r *http.Request) {
	mdFilePath := "." + r.URL.Path

	mdBytes, err := readMarkdownFile(mdFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	htmlBytes := convertMarkdownToHTML(mdBytes)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
<style>
  pre {
    border: 1px solid #ccc;
    padding: 1em;
    background-color: #f8f8f8;
  }
  code {
    white-space: pre-wrap;
  }
</style>
</head>
<body>
`)); err != nil {
		log.Println("Error writing to response:", err)
	}
	if _, err := w.Write(htmlBytes); err != nil {
		log.Println("Error writing to response:", err)
	}
	if _, err := w.Write([]byte(`</body>
</html>
`)); err != nil {
		log.Println("Error writing to response:", err)
	}
}

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/md/", serveMarkdownFile)
	fmt.Println("Listening on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe failed:", err)
	}
}
