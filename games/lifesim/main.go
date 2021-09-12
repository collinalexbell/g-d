package lifesim

import (
	"fmt"
	"io/ioutil"
	"os"
)

type WorldChunk = [][]string


func Run() {
	fmt.Println("welcome to the life sim!")
	os.Chdir("games/lifesim")
	bedroom_data, err := ioutil.ReadFile("kuberlogs_bedroom")
	if(err != nil) {
		fmt.Println("error reading bedroom_data: " + err.Error())
		return
	}
	fmt.Println(string(bedroom_data))
}

