package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var ExtSlice = []string{"doc", "docx", "html", "jpeg", "jpg", "JPG", "pdf", "pptx", "xlsx", "zip", "mp3", "mp4", "sh", "txt", "py", "ipynb", "csv", "conf", "png", "xls", "part", "PNG", "xls", "rar", "gz", "tar.gz", "xml", "ttf"}

func CheckAndCreateFolder(basedir string) ([]string, error) {

	folderDirs := []string{}

	// Creates folders if none exists for the extensions above...
	for _, ext := range ExtSlice {

		folder := filepath.Join(basedir, ext+"folder")
		folderDirs = append(folderDirs, folder)

		err := os.MkdirAll(folder, 0755)

		if err != nil {

			return nil, fmt.Errorf("failed to create directory %s: %s", folder, err)

		}

	}

	return folderDirs, nil

}

func GetFilesInFolder(path string) []fs.DirEntry {

	// Open and read path
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

func IsValidExtension(ext string) bool {
	for _, validExt := range ExtSlice {
		if ext == validExt {
			return true
		}
	}

	return false
}

func copyFile(sourcePath, destPath string) error {
	sourceFile, err := os.Open(sourcePath)

	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", sourcePath, err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create(destPath)

	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", destPath, err)
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)

	if err != nil {
		return fmt.Errorf("failed to copy file from %s to %s: %w", sourcePath, destPath, err)
	}

	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		return fmt.Errorf("destination file %s does not exist after copy", destPath)
	}

	return nil
}

func moveFile(sourcePath, destPath string) error {

	err := os.Rename(sourcePath, destPath)

	if err != nil {
		err = copyFile(sourcePath, destPath)

		if err != nil {
			return err
		}
		err = os.Remove(sourcePath)

		if err != nil {
			return fmt.Errorf("failed to remove source file %s: %w", sourcePath, err)
		}
	}

	return nil
}

func EnsureDirExists(dirPath string) error {

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {

		err := os.MkdirAll(dirPath, 0755)

		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}

	}

	return nil
}

func checkFolderContainingExt(extension string, destDirs []string) (folder string) {

	for _, folder = range destDirs {

		if strings.Contains(folder, extension) {
			return folder
		}

	}

	return ""

}

func MoveFiles(baseDir string, files []fs.DirEntry, destDir []string) error {

	// Match up file names with the destination dirs they should be moved to

	for _, file := range files {

		fileName := file.Name()

		ext := getFileExtension(fileName)

		if IsValidExtension(ext) {

			folderContainingExt := checkFolderContainingExt(ext, destDir)

			if folderContainingExt != "" {

				sourcePath := filepath.Join(baseDir, fileName)

				destPath := filepath.Join(folderContainingExt, fileName)

				err := moveFile(sourcePath, destPath)

				if err != nil {
					return fmt.Errorf("failed to move file %s to %s: %w", sourcePath, destPath, err)
				}

			}

		}

	}

	return nil
}
