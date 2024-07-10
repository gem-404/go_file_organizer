package main

const Folder = "/home/gem-404/Downloads"

func main() {

	CheckAndCreateFolder()

	files := GetFilesInFolder(Folder)

	CheckOutFiles(files)

}
