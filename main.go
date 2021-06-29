package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, fileInfo os.FileInfo, err error) error {
		fmt.Println(path)

		return nil
	})
}