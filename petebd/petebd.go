package petebd

import (
	"fmt"
	"github.com/gdamore/tcell"
)

var (
	screen tcell.Screen
)

var styleDefault = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

func allInit() {

	fmt.Println("PRO TIP: It is recommended that you FULLY MAXIMIZE this window\n")

	terrainInit()
	charactersInit()
}

type MessageDoQuit struct{}

type MessageDoRender struct{ Event tcell.Event }

func Start() {
	allInit()

	screen, _ = tcell.NewScreen()
	screen.Init()

	doRender := make(chan MessageDoRender)
	doQuit := make(chan MessageDoQuit)

	go eventHandler(doRender, doQuit)
	go characterAIHandler(doRender, doQuit)

GameLoop:
	for {
		select {
		case <-doQuit:
			break GameLoop
		case <-doRender:
			screen.SetStyle(styleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
			screen.Clear()
			renderMap(screen)
			renderCharacters(screen)
			renderConversation(screen)
			writeString(screen, getWindowWidth()-20, getWindowHeight()-1, styleDefault.Foreground(tcell.ColorGray), "press q to quit.")
			screen.Show()
		}
	}

	screen.Fini()
}
