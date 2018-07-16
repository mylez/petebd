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

var randomQuotes = []string{
	"Get busy living or get busy dying.",
	"Twenty years from now you will be more disappointed by the things that you didn't do than by the ones you did do.",
	"Without an open-minded mind, you can never be a great success.",
	"Maybe ever'body in the whole damn world is scared of each other.",
	"Some men get the world, some men get ex-hookers and a trip to Arizona. You're in with the former, but my God I don't envy the blood on your conscience.",
	"It is a far, far better thing that I do, than I have ever done; it is a far, far better rest I go to than I have ever known.",
	"But soft! What light through yonder window breaks? It is the east, and Juliet is the sun.",
	"I was at school with Tony Blair's brother Lionel",
	"Aren't all cars time machines?",
	"You look great today.",
	"You're a smart cookie.",
	"I bet you make babies smile",
	"You have impeccable manners.",
	"I like your style.",
	"You have the best laugh.",
	"I appreciate you.",
	"My, my, my!",
	"I really like the new landscaping job.",
	"Does anyone want to play with the blocks?",
	"Pete has a trampoline in that room. I am in line to try it.",
	"Is everyone ready for ice cream and pie?",
	"I think that chocolate and peanut butter mixed together is weird.",
	"Does anyone know where my glass went?",
	"When will the lamb be ready?",
	"I can't find my iPhone.",
	"Has anyone ever told you that you have great posture?",
	"How do you keep being so funny and making everyone laugh?",
	"Babies and small animals probably love you.",
	"I bet you do the crossword puzzle in ink.",
	"What do you think is the most important political issue at the moment?",
	"Do you enjoy debating politics with your friends? Do they have similar views to yours?",
	"Do you think people's political views change over their lifetime?",
	"Do you think too much money is spent on political campaigns?",
	"In the US and sometimes in the UK actors or actresses get involved in politics. Does this happen in your country?",
	"It is sometimes claimed that there is little difference between nationalism and racism. What is your opinion?",
	"Should a person be required to prove citizenship prior to casting their vote?",
	"Hey, when someone moves from one place to another, everything in-between is also called a place, right?",
	"Can I bum a cigarette?",
	"Tell me about your mother.",
	"What do you think of when I say the word Cerulean?",
	"Have you done any short weekend getaways recently? I'd love some recommendations.",
	"Tell me three unlikely things you did today.",
	"Where's the least attractive place to have hair? I'd say it's the tongue.",
	"I'll bet you're from a foreign country.",
	"I can't believe Harry Potter only got with two girls in school. Poor lad.",
	"I'm dying for a pack of Smarties.",
	"I'm genuinely so upset that Sabrina The Teenage Witch isn't on TV anymore. Aren't you?!",
	"Imagine you had amnesia and didn't know it? Hi, by the way.",
	"Imagine everyone in here was suddenly a cartoon? How would you even deal?",
	"Whenever something becomes popular I generally run in the opposite direction, it's my bias.",
	"That seems a little creepy to me...",
	"Happy birthday, Pete!",
	"Happy birthday, Pete!",
	"Happy birthday, Pete!",
	"Happy birthday, Pete!",
	"Happy birthday, Pete!",
}

var randomRages = []string{
	"Hey, watch out!",
	"Watch it!",
	"Do you not see me standing here?",
	"Ouch!",
	"Hmph!",
	"Oof!",
	"You almost spilled my drink!",
	"Be more careful.",
	"Excuse you.",
	"Hey, I'm walkin' here! I'm walkin' here!",
}

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
		Style:     styleDefault.Foreground(c.Color),
	})
}
func debugSays(msg string) {
	conversation = append(conversation, statement{
		Message:   "DEBUG: " + msg,
		Timestamp: time.Now().Format("3:04:05 PM -"),
		Style:     styleDefault,
	})
}
