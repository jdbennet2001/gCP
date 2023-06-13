package fsutils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

func Mirror( source string, target string, sizeBytes int64) error{

	//Get a list of all files in the source directory, above a given size
	fmt.Println("Loading file list")
	list := Walk(source, sizeBytes)

	for index, sourcePath := range list {

		relativePath, _ := getRelativePath(source, sourcePath)
		targetPath := filepath.Join(target, relativePath)

		tgtExists, _ := fileExists(targetPath)
		sameSize, _ := areFilesSameSize(sourcePath, targetPath)

		progress := fmt.Sprintf("%d / %d", index, len(list))

		if tgtExists && sameSize{
			fmt.Println(".. ", progress, " skipping: ", sourcePath, ", target exists")
		}else{
			fmt.Println(".. ", progress, " copying:  ", sourcePath, " --> ", targetPath)
			copy(sourcePath, targetPath)
		}

	}

	return nil
}

/*
 Move an archive to the target directory
*/
func copy(sourcePath string, targetPath string) {

	os.MkdirAll(path.Dir(targetPath), 0700)

	from, err := os.Open(sourcePath)
	if err != nil {
		log.Fatal(err)
	}

	to, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

}


func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // File does not exist
		}
		return false, err // Error occurred while checking file
	}
	return true, nil // File exists
}

func areFilesSameSize(file1, file2 string) (bool, error) {
	info1, err := os.Stat(file1)
	if err != nil {
		return false, err
	}

	info2, err := os.Stat(file2)
	if err != nil {
		return false, err
	}

	return info1.Size() == info2.Size(), nil
}

func getRelativePath(directory, filename string) (string, error) {
	absDirectory, err := filepath.Abs(directory)
	if err != nil {
		return "", err
	}

	absFilename, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}

	relativePath, err := filepath.Rel(absDirectory, absFilename)
	if err != nil {
		return "", err
	}

	return relativePath, nil
}