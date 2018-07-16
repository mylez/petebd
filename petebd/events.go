package petebd

import (
	"github.com/gdamore/tcell"
)

func eventHandler(doRender chan MessageDoRender, doQuit chan MessageDoQuit) {

Loop:
	for {
		ev := screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				doQuit <- MessageDoQuit{}
				break Loop
			}

			playerHandleKeyEvent(ev)
			doRender <- MessageDoRender{ev}
		default:
			doRender <- MessageDoRender{ev}
		}
	}
}
