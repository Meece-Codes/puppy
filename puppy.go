package puppy

import (
	"fmt"

	"github.com/Meece-Codes/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func Barking() string {
	return Bark()
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func from13() {
	fmt.Println("I'm from version 1.3.0")
}
