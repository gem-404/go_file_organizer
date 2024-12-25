// Description: This is the main file of the program. It takes a folder path as
// an argument and then proceeds to organize the files in the folder by their
// extensions. It first checks if the folder exists, then gets the files in the
// folder, and then gets the extensions of the files. It then creates folders
// for each extension if they don't exist and moves the files to their
// respective folders.

package main

import (
	"log"
	"os"
)

func main() {

	// Check if the user provided the folder path as an argument
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <folder_path>", os.Args[0])
	}

	baseDir := os.Args[1]

	err := EnsureDirExists(baseDir)

	if err != nil {
		log.Fatalf("Directory %s does not exist:%s\n", baseDir, err)
	}

	files := GetFilesInFolder(baseDir)

	extensions := GetExtensions(files)

	folderDirs, err := CheckAndCreateFolder(baseDir, extensions)

	if err != nil {
		log.Fatalf("An error occurred during creating folders in baseDir: %s, %s", baseDir, err)
	}

	err = MoveFiles(baseDir, files, folderDirs)

	if err != nil {
		log.Fatalf("An error occurred while moving files: %s", err)
	}
}
