package files

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// A list of directories that shouldn't be displayed or written
var bannedDirectories = []string{
	// Common
	".git",

	// Linux system directories
	"bin", "boot", "dev", "etc", "lib", "lib32", "lib64", "libx32",
	"lost+found", "opt", "proc", "root", "run", "sbin",
	"snap", "srv", "sys", "tmp", "usr", "var",

	// Windows system directories
	"Windows", "System32", "SysWOW64", "WinSxS", "Program Files",
	"Program Files (x86)", "ProgramData", "PerfLogs", "Recovery",
	"$Recycle.Bin", "$WinREAgent", "System Volume Information",
	"Documents and Settings", "MSOCache", "OneDriveTemp",
}

// File structure of how to write to json output
type File struct {
	path string
	size float64
}

// Temp code storage to write them to json
type FileStorage struct {
	database []File
}

// Method to add files to file database to later write to json
func (db *FileStorage) addToDatabase(file File) {
	db.database = append(db.database, file)
}

// Returns a string that contains the value of megabytes in a file
func (f File) SizeString() string {
	return strconv.FormatFloat(f.size, 'f', -1, 64) + "mb"
}

// Declaring temp file database
var FileDatabase = FileStorage{}


func ReadDirectory() {
	items, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}


	var dirs []string
	for _, entry := range items {
		if entry.IsDir() {
			isBanned := false
			for _, banned := range bannedDirectories {
				if banned == entry.Name() {
					isBanned = true
					break
				}
			}
			if !isBanned {
				dirs = append(dirs, entry.Name())
			}
		} else {
			info, err := os.Stat(entry.Name())
			if err != nil {
				continue
			}
			file := File{
				path: entry.Name(),
				size: float64(info.Size()) / (1024 * 1024),
			}
			fmt.Println(file.path + ": " + file.SizeString())
		}
	}

	for _, dir := range dirs {
		prev, err := os.Getwd()
		if err != nil {
			continue
		}
		if err := os.Chdir(dir); err != nil {
			continue
		}
		ReadDirectory()
		os.Chdir(prev)
	}
}
