package main

import (
	"fmt"
	"mass/files"
	"mass/utils"
)


func main() {
	fmt.Println(utils.InitialDirectory)
	files.ReadCurrentDirectory()
}
