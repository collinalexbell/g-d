package main


import (
	"os/user"
	"fmt"
)


func reorganizeFactory() {
	zones := []string{
		"kitchen",
		"bedroom",
		"closet",
		"laundry room",
		"hallway",
		"bathroom",
		"sunroom",

		// Rerogranize these zones
	}
	for _, zone := range zones {
		fmt.Println("Reorganizing", zone)
	}
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s's factorio IRL regorganizer", user)
		reorganizeFactory()
	}
}


