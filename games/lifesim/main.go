package lifesim

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"encoding/json"
	"github.com/rthornton128/goncurses"
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

func scan_worldchunk(worldchunk_filename string) World_Chunk {
	legend_json := ""
	var world_chunk World_Chunk

	f, err := os.Open(worldchunk_filename)
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
			world_chunk.Map = append(world_chunk.Map, make([]rune, 0))
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
	return world_chunk
}

func game_loop(screen *goncurses.Window, world_chunk *World_Chunk) {
	pos := []int{5,5}
	var tmp rune
	for {
		tmp = world_chunk.Map[pos[0]][pos[1]]
		world_chunk.Map[pos[0]][pos[1]] = '&'
		screen.Clear()
		for _, y_data := range(world_chunk.Map) {
			for _, r := range(y_data) {
				screen.AddChar(goncurses.Char(r))
			}
			screen.Println()
		}
		key := screen.GetChar()

		// 
		world_chunk.Map[pos[0]][pos[1]] = tmp
		switch byte(key) {
		case 'q':
			return
		case 'w':
			pos = []int{pos[0]-1, pos[1]}
		case 's':
			pos = []int{pos[0]+1, pos[1]}
		case 'd':
			pos = []int{pos[0], pos[1]+1}
		case 'a':
			pos = []int{pos[0], pos[1]-1}
		}
		screen.Refresh()

	}
}


func Run() {
	screen, err := goncurses.Init()
	if err != nil {
		fmt.Println("ncurses init failed")
		return
	}
	defer goncurses.End()
	fmt.Println("welcome to the life sim!")
	os.Chdir("games/lifesim")
	world_chunk := scan_worldchunk("kuberlogs_studio")
	game_loop(screen, &world_chunk)

}

