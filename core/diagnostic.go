package core

/*
 * Diagnostic is a part of error
 * with the source code location and label provided
 */
type Diagnostic struct {
	SourcePath string
	Source     *string
	Location   Location
	Label      string
}

// Creates new diagnostic
func NewDiagnostic(source string, location Location, label string) Diagnostic {
	return Diagnostic{
		SourcePath: source,
		Location:   location,
		Label:      label,
	}
}

func NewDiagnosticWithSource(path, source string, location Location, label string) Diagnostic {
	return Diagnostic{
		SourcePath: path,
		Source:     &source,
		Location:   location,
		Label:      label,
	}
}
