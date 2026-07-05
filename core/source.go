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
func ReadSourceDiagnostic(diagnostic Diagnostic, style *Style) (string, string) {
	if diagnostic.Source != nil {
		return filepath.Base(diagnostic.SourcePath), *diagnostic.Source
	}

	// Openning source code file
	file, err := os.Open(diagnostic.SourcePath)
	if err != nil {
		fmt.Println(NewError(fmt.Sprintf("Could not read source code file: %s", diagnostic.SourcePath)).Error(style))
		os.Exit(1)
	}
	defer file.Close()

	// Reading file name
	fileName := filepath.Base(diagnostic.SourcePath)

	// Reading source code
	scanner := bufio.NewScanner(file)
	n := 0
	for scanner.Scan() {
		if n < diagnostic.Location.Line {
			n++
		} else {
			return fileName, scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(NewError(fmt.Sprintf("Error reading source code file: %s", diagnostic.SourcePath)).Error(style))
		os.Exit(1)
	}

	return fileName, scanner.Text()
}
