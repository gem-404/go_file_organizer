# File Organizer in Go

This project is a simple file organizer written in Go. It reads a specified directory, identifies files by their extensions, and moves them into corresponding folders based on these extensions. The folders are created dynamically based on the extensions of the files present in the directory.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Functions](#functions)
  - [Main Functions](#main-functions)
  - [Helper Functions](#helper-functions)
- [Contributing](#contributing)
- [License](#license)

## Features

- Automatically identifies and organizes files based on their extensions.
- Creates necessary folders dynamically.
- Handles files with various common extensions.
- Ensures no empty folders are created.

## Requirements

- Go 1.16 or later

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/gem-404/file-organizer-go
    cd file-organizer-go
    ```

2. **Build the project:**

    ```sh
    go build .
    ```

## Usage

To run the file organizer, provide the path to the directory you want to organize as an argument:

```sh
go run . /path/to/your/folder
```

For example:

```sh
go run . /home/gem-404/Downloads
```

## Project Structure

```plaintext
.
├── main.go
├── file_organizer.go
└── README.md
```

- **main.go**: Contains the main function and argument handling.
- **file_organizer.go**: Contains all the helper functions and core logic for organizing files.
- **README.md**: Documentation for the project.

## Functions

### Main Functions

- **main()**:
  - Checks for the directory argument.
  - Calls functions to ensure directory existence, create necessary folders, get files, and move files.

### Helper Functions

- **EnsureDirExists(dirPath string) error**:
  - Ensures that a directory exists, creating it if necessary.

- **CheckAndCreateFolder(baseDir string, extensions []string) ([]string, error)**:
  - Creates folders based on the file extensions present in the directory.

- **GetFilesInFolder(path string) []fs.DirEntry**:
  - Reads the directory and returns a slice of directory entries (files only).

- **getFileExtension(fileName string) string**:
  - Extracts the file extension from a given file name.

- **IsValidExtension(ext string) bool**:
  - Checks if a given extension is in the predefined list of extensions.

- **GetExtensions(files []fs.DirEntry) []string**:
  - Retrieves a list of unique file extensions from the provided files.

- **checkFolderContainingExt(extension string, destDirs []string) string**:
  - Checks which folder contains a specific extension.

- **copyFile(sourcePath, destPath string) error**:
  - Copies a file from source to destination.

- **moveFile(sourcePath, destPath string) error**:
  - Moves a file from source to destination, using copying if necessary.

- **MoveFiles(baseDir string, files []fs.DirEntry, destDirs []string) error**:
  - Moves files to their respective folders based on their extensions.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

By following the steps in this README, you should be able to successfully use and understand the file organizer project. If you encounter any issues or have any questions, feel free to open an issue or submit a pull request.
