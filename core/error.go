package core

/*
 * Represents error
 * as diagnostics sequence
 */
type Error struct {
	Title       string
	Diagnostics []Diagnostic
}

// Creates new error
func NewError(title string) *Error {
	return &Error{
		Title:       title,
		Diagnostics: nil,
	}
}

// With diagnostics
func (self *Error) WithDiagnostics(diagnostics []Diagnostic) *Error {
	self.Diagnostics = append(self.Diagnostics, diagnostics...)
	return self
}

// With diagnostic
func (self *Error) WithDiagnostic(diagnostic Diagnostic) *Error {
	self.Diagnostics = append(self.Diagnostics, diagnostic)
	return self
}

// Diagnostic to string
// implementation of the error interface
func (self *Error) Error(style *Style) string {
	return Report(self, style)
}
