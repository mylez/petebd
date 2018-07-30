package petebd

import (
	"fmt"
	"github.com/gdamore/tcell"
	"math/rand"
	"time"
)

type statement struct {
	Character *character
	Message   string
	Timestamp string
	Style     tcell.Style
}

var conversation []statement

func randomRage() string {
	return randomRages[rand.Int()%len(randomRages)]
}

func randomQuote() string {
	return randomQuotes[rand.Int()%len(randomQuotes)]
}

func renderConversation(screen tcell.Screen) {
	l := 0
	for j := getWindowHeight() - 1; j > mapHeight && l < len(conversation); j-- {
		c := conversation[len(conversation)-l-1]
		writeString(screen, 0, j, c.Style, fmt.Sprintf("%s %s\n", c.Timestamp, c.Message))
		l++
	}
}

func alert(msg string) {
	conversation = append(conversation, statement{
		Message: msg, Style: styleDefault.Foreground(tcell.ColorRed),
		Timestamp: time.Now().Format("3:04:05 PM -"),
	})
}

func somebodyBumps(a, b *character) {
	conversation = append(conversation, statement{
		Message:   fmt.Sprintf("%s bumps into %s", a.Name, b.Name),
		Timestamp: time.Now().Format("3:04:05 PM -"),
		Style:     styleDefault.Foreground(tcell.ColorGreen),
	})

	if rand.Int()%2 == 0 {
		somebodySays(b, randomRage())
	}
}

func somebodySays(c *character, message string) {
	conversation = append(conversation, statement{
		Message:   fmt.Sprintf("%s says: %s", c.Name, message),
		Timestamp: time.Now().Format("3:04:05 PM -"),
		Style:     c.Style,
	})
}
func debugSays(msg string) {
	conversation = append(conversation, statement{
		Message:   "DEBUG: " + msg,
		Timestamp: time.Now().Format("3:04:05 PM -"),
		Style:     styleDefault,
	})
}
