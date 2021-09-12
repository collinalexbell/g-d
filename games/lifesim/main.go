package lifesim

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"encoding/json"
)

type World_Chunk struct {
	Legend map[string]string
	Map [][]rune
	Description string
}

func get_section(chunk_scanner *bufio.Scanner, section *string) {
	if strings.Contains(chunk_scanner.Text(), "</") {
		*section = ""
	}
	if strings.Contains(chunk_scanner.Text(), "<LEGEND>") {
		*section = "legend"
	}
	if strings.Contains(chunk_scanner.Text(), "<MAP>") {
		*section = "map"
	}
	if strings.Contains(chunk_scanner.Text(), "<DESCRIPTION>") {
		*section = "description"
	}
}

func scan_worldchunk(worldchunk_filename string) {
	legend_json := ""
	var world_chunk World_Chunk

	f, err := os.Open("kuberlogs_bedroom")
	if err != nil {
		fmt.Println("failed to open kuberlogs_bedroom: " + err.Error())
	}
	defer f.Close()
	chunk_scanner := bufio.NewScanner(f)
	last_section := ""
	section := ""

	for chunk_scanner.Scan() {
		get_section(chunk_scanner, &section)
		if section != last_section && section != "" {
			fmt.Println("found " + section)
			last_section = section
			continue
		}
		if section != last_section && section == "" {
			fmt.Println("leaving " + last_section)
			last_section = section
			continue
		}

		switch section {
		case "legend":
			legend_json += chunk_scanner.Text()
		case "map":
			line := chunk_scanner.Text()
			world_chunk.Map = append(world_chunk.Map, make([]rune, len(line)))
			y := len(world_chunk.Map)-1
			for _, r := range(line) {
				world_chunk.Map[y] = append(world_chunk.Map[y], r)
			}
		case "description":
			world_chunk.Description += chunk_scanner.Text()
		}
	}

	err = json.Unmarshal([]byte(legend_json), &(world_chunk.Legend))
	if err != nil {
		fmt.Println("error parsing the legend")
	}
	fmt.Println(world_chunk)
}


func Run() {
	fmt.Println("welcome to the life sim!")
	os.Chdir("games/lifesim")
	scan_worldchunk("kuberlogs_bedroom")

}

