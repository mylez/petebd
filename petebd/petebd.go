package petebd

import (
	"fmt"
	"github.com/gdamore/tcell"
	"time"
)

var screen tcell.Screen

var styleDefault = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

func allInit() {
	fmt.Println("PRO TIP: It is recommended that you FULLY MAXIMIZE this window")
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
	go blinkOnStartup(doRender, doQuit)

GameLoop:
	for {
		select {
		case <-doQuit:
			break GameLoop
		case <-doRender:
			screen.SetStyle(styleDefault)
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

func blinkOnStartup(doRender chan MessageDoRender, doQuit chan MessageDoQuit) {
	for i := 0; i < 15; i++ {
		time.Sleep(300 * time.Millisecond)
		peteCharacter.Style = styleDefault.Background(tcell.ColorYellow).Foreground(tcell.ColorBlack)
		doRender <- MessageDoRender{}
		time.Sleep(300 * time.Millisecond)
		peteCharacter.Style = styleDefault
		doRender <- MessageDoRender{}
	}
}
