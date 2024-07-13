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
- [Building](#Building)
- [Usage](#Usage)
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


## Building

To build the executable and make it accessible system-wide, follow these steps:

1. **Build the Go Program:**

   Navigate to your project directory and build the binary.

   ```sh
   cd /path/to/your/project
   go build -o fileorganizer
   ```

   This command will generate an executable named `fileorganizer` in your project directory.

2. **Move the Executable to a Directory in Your PATH:**

   Move the executable to `/usr/local/bin` or any directory that's included in your `PATH`.

   ```sh
   sudo mv fileorganizer /usr/local/bin/
   ```

3. **Verify the Executable:**

   Ensure that the executable is accessible from anywhere by checking the `PATH`.

   ```sh
   echo $PATH
   which fileorganizer
   ```

   This should show `/usr/local/bin/fileorganizer` or the path where you moved the executable.

## Usage

To run the program from any directory, use the following command:

```sh
fileorganizer /path/to/target/directory
```

## Adding the Program to PATH Manually (Optional)

If you prefer to add the directory containing your executable to the `PATH` manually, follow these steps:

1. **Edit the Shell Configuration File:**

   Open your preferred text editor and edit the appropriate configuration file (`.bashrc`, `.zshrc`, or `.profile`).

   ```sh
   nano ~/.bashrc
   ```

   or

   ```sh
   nano ~/.zshrc
   ```

2. **Add the Directory to PATH:**

   Add the following line to the file:

   ```sh
   export PATH=$PATH:/path/to/your/executable
   ```

   For example:

   ```sh
   export PATH=$PATH:/home/yourusername/go/bin
   ```

3. **Reload the Configuration:**

   After saving the file, reload the configuration.

   ```sh
   source ~/.bashrc
   ```

   or

   ```sh
   source ~/.zshrc
   ```

Now, your Go program is accessible system-wide, and you can execute it from any directory without needing to navigate to the project directory. This approach simplifies running your utility, especially when automating tasks or integrating them into scripts.

## Example Usage

Here’s an example of how to use the `fileorganizer` program:

```sh
fileorganizer /home/user/Downloads
```

This will organize all files in the `/home/user/Downloads` directory by moving them into subdirectories based on their file extensions.


## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

By following the steps in this README, you should be able to successfully use and understand the file organizer project. If you encounter any issues or have any questions, feel free to open an issue or submit a pull request.
