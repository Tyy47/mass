package files

import (
	"fmt"
	"log"
	"os"
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


var FileDatabase = FileStorage{}


func ReadCurrentDirectory() {
	item, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	for _, entry := range item {
		fmt.Println(entry)
	}

	fmt.Println("\n Function complete")
}
