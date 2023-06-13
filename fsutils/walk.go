package fsutils

import (
"fmt"
"os"
"path/filepath"
"strings"
)

//Walk a directory tree returning a of files with a given extension
func Walk(rootpath string, sizeBytes int64)[]string {

	var results []string

	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {

		// Apple cache data is never useful
		basename := filepath.Base(path)
		if strings.HasPrefix(basename, "."){
			return nil
		}

		large, err :=  isFileOverSize(path, sizeBytes)

		if err != nil{
			fmt.Printf("FS Size error [%v]\n", err)
			return nil
		}

		if large{
			results = append(results, path)
		}

		return nil	// No error...
	})

	if err != nil {
		fmt.Printf("Walk error [%v]\n", err)
	}

	return results

}

func isFileOverSize(filePath string, sizeBytes int64) (bool, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}

	fileSize := fileInfo.Size()
	return fileSize > sizeBytes, nil
}
