package files

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type File struct {
	path string
	size float64
}

type FileStorage struct {
	database []File
}

func (db *FileStorage) addToDatabase(file File) {
	db.database = append(db.database, file)
}

func (f File) SizeString() string {
	return strconv.FormatFloat(f.size, 'f', -1, 64) + "mb"
}


var FileDatabase = FileStorage{}


func ReadCurrentDirectory() {
	items, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}


	var dirs []string
	for _, entry := range items {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		} else {
			info, _ := os.Stat(entry.Name())
			file := File{
				path: entry.Name(),
				size: float64(info.Size()) / (1024 * 1024),
			}
			fmt.Println(file.path + ": " + file.SizeString())
		}
	}

	for _, dir := range dirs {
		os.Chdir(dir)
		ReadCurrentDirectory()
		os.Chdir("..")
	}
}
