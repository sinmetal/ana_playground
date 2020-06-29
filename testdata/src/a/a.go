package a

import (
	"fmt"     // want "fmt is duplicated import"
	fff "fmt" // want "fff fmt is duplicated import"
)

func f() {
	// The pattern can be written in regular expression.
	var gopher int // want "identifyer is gopher"
	print(gopher)  // want "identifyer is gopher"

	_ = fmt.Sprintf("")
	_ = fff.Sprintf("")
}
