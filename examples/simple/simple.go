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
		))
	fmt.Printf("%s", error.Error(c.DefaultStyle()))
}
