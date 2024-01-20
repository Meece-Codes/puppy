package puppy

import (
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
