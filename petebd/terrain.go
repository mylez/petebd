package petebd

import (
	"github.com/gdamore/tcell"
	"math/rand"
	"strings"
	"time"
)

var (
	mapData   [][]rune
	mapWidth  int = 0
	mapHeight int = 0
)

func terrainInit() {
	loadMap()

	rand.Seed(time.Now().Unix())
}

func renderMap(screen tcell.Screen) {
	w, h := screen.Size()

	if w > mapWidth {
		w = mapWidth
	}

	if h > mapHeight {
		h = mapHeight
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			r := getMapTile(x, y)
			screen.SetCell(x, y, styleForTile(r), r)
		}
	}
}

func isWalkable(x, y int) bool {
	mt := getMapTile(x, y)

	for _, npc := range nonPlayableCharacters {
		if npc.PosX == x && npc.PosY == y {
			return false
		}
	}

	return mt == ' ' || mt == 'D' || mt == 'H'
}

func getMapTile(x, y int) rune {
	if y < 0 || x < 0 {
		return '#'
	}
	if y < len(mapData) {
		if x < len(mapData[y]) {
			return mapData[y][x]
		}
	}
	return '#'
}

func loadMap() {
	lines := strings.Split(mapRawText, "\n")
	mapWidth = 0
	mapHeight = 0

	for _, line := range lines {
		mapHeight++
		l := len(line)

		if l > mapWidth {
			mapWidth = l
		}
	}

	mapData = makeRuneMatrix(mapWidth, mapHeight, ' ')

	for y, line := range lines {
		for x, r := range line {
			mapData[y][x] = r
		}
	}
}

var defibFlash = 0

func styleForTile(tile rune) tcell.Style {
	switch tile {
	case 'P':
		// piano
		return styleDefault.Foreground(tcell.NewRGBColor(46, 46, 46))
	case 'S':
		// sink
		return styleDefault.Foreground(tcell.ColorDarkBlue)
	case 'T':
		// table
		return styleDefault.Foreground(tcell.ColorSaddleBrown)
	case 'C':
		// counter
		return styleDefault.Foreground(tcell.ColorDarkOliveGreen)
	case 'K':
		// kitchen counter
		return styleDefault.Foreground(tcell.NewRGBColor(50, 40, 90))
	case 'F':
		// fridge
		return styleDefault.Foreground(tcell.NewRGBColor(200, 200, 200))
	case '#':
		// wall
		return styleDefault.Foreground(tcell.NewRGBColor(90, 90, 90))
	case '~':
		// defibrillator
        return styleDefault.Foreground(tcell.ColorRed)
	case '&':
		// stuffed animal
		return styleDefault.Foreground(tcell.ColorBrown)
	default:
		return styleDefault
	}
}
