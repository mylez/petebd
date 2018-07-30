package petebd

import "github.com/gdamore/tcell"

func makeRuneMatrix(w, h int, initial rune) [][]rune {
	var m [][]rune

	for y := 0; y < h; y++ {
		m = append(m, []rune{})
		for x := 0; x < w; x++ {
			m[y] = append(m[y], initial)
		}
	}

	return m
}

func writeString(screen tcell.Screen, x, y int, s tcell.Style, msg string) {
	for i, r := range msg {
		screen.SetCell(x+i, y, s, r)
	}
}

func getWindowHeight() int {
	_, h := screen.Size()
	return h
}

func getWindowWidth() int {
	w, _ := screen.Size()
	return w
}
