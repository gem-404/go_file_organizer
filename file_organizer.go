package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var ExtSlice = []string{"doc", "docx", "html", "jpeg", "jpg", "JPG", "pdf", "pptx", "xlsx", "zip", "mp3", "mp4", "sh", "txt", "py", "ipynb", "csv", "conf", "png", "xls", "part", "PNG", "xls", "rar", "gz", "tar.gz"}

func CheckAndCreateFolder() {

	// Creates folders if none exists for the extensions above...
	for _, ext := range ExtSlice {

		folder := ext + "folder"

		err := os.MkdirAll(folder, 0755)

		if err != nil {

			log.Fatalf("failed to create directory %s: %s", folder, err)

		}

	}

}

func GetFilesInFolder(path string) []fs.DirEntry {

	// Open path
	dirEntries, err := os.ReadDir(path)

	if err != nil {
		log.Fatalf("failed to read directory %s: %s", path, err)
	}

	var files []fs.DirEntry

	for _, entry := range dirEntries {

		if !entry.IsDir() {
			files = append(files, entry)
		}

	}

	return files
}

func getFileExtension(fileName string) string {
	ext := filepath.Ext(fileName)
	return strings.TrimPrefix(ext, ".") // Remove the leading dot
}

func isValidExtension(ext string) bool {
	for _, extension := range ExtSlice {
		if ext == extension {
			return true
		}
	}

	return false
}

func CheckOutFiles(files []fs.DirEntry) {

	for _, file := range files {
		// Remove single quotes from the file name
		fileName := strings.ReplaceAll(file.Name(), "'", "")

		// Extract the file extension
		ext := getFileExtension(fileName)

		// Check if the file extension is in the valid list
		if isValidExtension(ext) {
			// fmt.Printf("Valid file found: %s\n", fileName)
			fmt.Println(ext)
		} else {
			fmt.Printf("File ignored: %s\n", fileName)
		}
	}

}
