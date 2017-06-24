package main

import (
	"fmt"
	"os"
	"strconv"
)

func gatherPanics() {
	if rec := recover(); rec != nil {
		fmt.Println("Critical Error:", rec)
	}
}

func getFileInfo(filename string) {
	defer gatherPanics()
	finfo, err := os.Stat(filename)
	if err != nil {
		panic("File not valid!")
	} else {
		fmt.Println("Size:", strconv.FormatInt(finfo.Size(), 10))
	}
}

func openFile(filename string) {
	defer gatherPanics()
	if _, err := os.Stat(filename); err != nil {
		panic("File does not exist!")
	}
}

func main() {
	var filename string
	fmt.Println("Enter Filename:")
	_, err := fmt.Scanf("%s", &filename)
	if err != nil {
	}

	fmt.Println("Getting info for", filename)

	openFile(filename)
	getFileInfo(filename)

}
