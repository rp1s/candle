package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

/*
 * Reads source by path and location,
 * returns filename and line to render
 */
func ReadSource(path string, location Location, style *Style) (string, string) {
	// Openning source code file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(NewError(fmt.Sprintf("Could not read source code file: %s", path)).Error(style))
		os.Exit(1)
	}
	defer file.Close()

	// Reading file name
	fileName := filepath.Base(path)

	// Reading source code
	scanner := bufio.NewScanner(file)
	n := 0
	for scanner.Scan() {
		if n < location.Line {
			n++
		} else {
			return fileName, scanner.Text()
		}
	}

	return fileName, scanner.Text()
}
