package a

import fmt1 "fmt" // want "fmt1 fmt is duplicated import"
import fmt2 "fmt" // want "fmt2 fmt is duplicated import"

func b() {
	fmt1.Println("Hello")
	fmt2.Println("World")
}
