package main

import (
	"fmt"
	"log"
	"os"
)

// const baseDir = "/home/gem-404/Downloads"

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

	folderDirs, err := CheckAndCreateFolder(baseDir)

	if err != nil {
		log.Fatalf("An error occurred during creating baseDir: %s, %s", baseDir, err)
	} else {
		fmt.Println("All baseDirs created successfully!")
	}

	files := GetFilesInFolder(baseDir)

	err = MoveFiles(baseDir, files, folderDirs)

	if err != nil {
		log.Fatalf("An error occurred while moving files: %s", err)
	} else {
		fmt.Println("All files moved successfully!")
	}

}
