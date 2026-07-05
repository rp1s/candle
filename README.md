# 🕯️ Candle
Candle is simple and fancy diagnostic reporting framework.

# Features
- Easy creation of errors with zero, one or multiple diagnostics.
- Customizable error display styles.
- Supports specifying file, line, column, and span locations for errors.

# Installation
```bash
go get github.com/vyacheslavhere/candle
```

# Usage
Simple error with two diagnostics:
```go
package main

import (
	c "candle/core"
	"fmt"
)

func main() {
	error := c.NewError("Error occurred").
		WithDiagnostic(c.NewDiagnostic(
			"./test.err",
			c.NewLocation(0, 7, 28),
			"Сould not concatenate `str` and `int`.",
		)).
		WithDiagnostic(c.NewDiagnostic(
			"./test.err",
			c.NewLocation(1, 5, 10),
			"Invalid function name.",
		)).
		WithDiagnostic(c.NewDiagnosticWithSource(
			"./test.err",
			"fmt.Println(\"Hello, World!\")",
			c.NewLocation(1, 5, 12),
			"Invalid function name.",
		))
	fmt.Printf("%s", error.Error(c.DefaultStyle()))
}
```
<img alt="image" src="assets/simple.png" />

# Customizing style
You can easily make your own style for errors, using Style struct:
```go
type Style struct {
	ErrorPrefix        string // Prefix for the error message (e.g. "Error: ")
	ErrorPrefixColor   string // Color for the error prefix
	TitleColor         string // Color for the error title
	FileNameColor      string // Color for the filename
	FileNameColon      string // Character after the filename (e.g. ":")
	FileNameColonColor string // Color for the colon
	LineNumberColor    string // Color for the line number
	LineColor          string // Color for the line text
	SpaceChar          string // Character used for spacing
	ArrowChar          string // Character used for pointing to the error span (e.g. "^")
	ArrowCharColor     string // Color for the arrow
	CaptionColor       string // Color for the diagnostic message
}
```
