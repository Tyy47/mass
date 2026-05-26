package main

import (
	"mass/files"
	"os"
)


func main() {
	os.Chdir("/")
	files.ReadDirectory()
}
