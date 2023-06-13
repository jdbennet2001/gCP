package main

import (
	"flag"
	"fmt"
	"gCP/fsutils"
	"log"
	"os"
)

const min_file_size = 1026 * 64 //64 minimum file size

func main() {
	fmt.Println("gCP usage: 'gCP --from <dir> --to <dir>")

	fromPtr := flag.String("from", "", "Source file path")
	toPtr := flag.String("to", "", "Destination file path")

	flag.Parse()

	from := *fromPtr
	to := *toPtr

	if from == "" || to == "" {
		fmt.Println("Please provide both 'from' and 'to' flags.")
		return
	}

	fromExists, fromErr := directoryExists(from)

	if !fromExists {
		log.Fatal("Missing directory, specify command line values and try again")
	}

	if fromErr != nil {
		log.Fatal("Error accessing directories")
	}

	fsutils.Mirror(from, to, min_file_size)

}

func directoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // Directory does not exist
		}
		return false, err // Error occurred while checking directory
	}
	return true, nil // Directory exists
}
