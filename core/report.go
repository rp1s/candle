package core

import (
	"fmt"
	"strings"
)

/*
 * Reports diagnostics, if they given
 */
func ReportDiagnostics(sb *strings.Builder, err *Error, style *Style) *strings.Builder {
	if len(err.Diagnostics) == 0 {
		return sb
	}

	for _, diagnostic := range err.Diagnostics {
		file, line := ReadSource(diagnostic.Source, diagnostic.Location, style)

		sb.WriteString("\n")

		sb.WriteString(style.FileNameColor)
		sb.WriteString(file)
		sb.WriteString(style.FileNameColonColor)
		sb.WriteString(style.FileNameColon)
		sb.WriteString("\n")

		sb.WriteString(style.LineNumberColor)
		sb.WriteString(fmt.Sprint(diagnostic.Location.Line))
		sb.WriteString(" ")
		sb.WriteString(style.LineColor)
		sb.WriteString(line)
		sb.WriteString("\n")

		lineNumLen := len(fmt.Sprint(diagnostic.Location.Line))
		spacesAmount := diagnostic.Location.Start + lineNumLen
		arrowsAmount := diagnostic.Location.End - diagnostic.Location.Start

		sb.WriteString(strings.Repeat(style.SpaceChar, spacesAmount))
		sb.WriteString(style.ArrowCharColor)
		sb.WriteString(strings.Repeat(style.ArrowChar, arrowsAmount))
		sb.WriteString("\n")

		sb.WriteString(strings.Repeat(style.SpaceChar, spacesAmount))
		sb.WriteString(style.CaptionColor)
		sb.WriteString(diagnostic.Label)
		sb.WriteString(ResetColor)
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	return sb
}

/*
 * Reports an error
 */
func Report(error *Error, style *Style) string {
	var sb strings.Builder

	sb.WriteString(style.ErrorPrefixColor)
	sb.WriteString(style.ErrorPrefix)
	sb.WriteString(style.TitleColor)
	sb.WriteString(" ")
	sb.WriteString(error.Title)
	sb.WriteString("\n")

	return ReportDiagnostics(&sb, error, style).String()
}
