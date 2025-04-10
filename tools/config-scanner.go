package main

import (
    "bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GOOS=windows GOARCH=amd64 go build -o windows-scanner.exe main.go

 func Scan(info os.FileInfo, path string) {
	filer, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s. Can't open file %s.", err.Error(), path)
	}
	defer filer.Close()
	scanner := bufio.NewScanner(filer) // NewScanner uses ScanLines func which returns each line of text, stripped of any trailing end-of-line marker.
	for scanner.Scan() {
		line := scanner.Text()
		configFile := "kind: Config"
		if strings.Contains(line, configFile) {
			fmt.Printf("Determined to be a Kubeconfig file: %s.\n", path)
		}
	}
 }

// Confirms n is a regular file and not a directory
func checkFile(info os.FileInfo, path string  ) {
if info.Mode().IsRegular() {
	// fmt.Printf("%s is a regular file\n", info.Name())
	filer, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s. Can't open file %s.", err.Error(), path)
	}
	defer filer.Close()
}
}

func main() {

	fileName := flag.String("file", "", "Filename pattern we're looking for, default is kubeconfig")
	startPath := flag.String("path", ".", "Where shall we scan, default is root")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Usage: ./exe -file <filename> [-path /start/dir]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("Searching for %s starting from %s...\n", *fileName, *startPath)

	err := filepath.Walk(*startPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), *fileName) {
			fmt.Printf("A file with Kubeconfig in its name: %s \n", path)

			checkFile(info, path)
			Scan(info, path)
		} else {
			checkFile(info, path)
			Scan(info, path)
		}
		return nil
})
    if err != nil {
        println("Can't walk the path provided. Error", err.Error())
    }

}
