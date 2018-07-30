package petebd

import (
	"bufio"
	"fmt"
	"github.com/gdamore/tcell"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	peteCharacter         *character
	nonPlayableCharacters []*character
)

func printEllipsis() {
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(250*time.Millisecond)
	}
	fmt.Println()
}

func charactersInit() {
	peteCharacter = randomizedCharacter("Pete", "M")
	peteCharacter.Style = styleDefault

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Invite family and friends? y/n: ")
	text, _ := reader.ReadString('\n')

	invitedFamily := false

	if len(text) > 0 && strings.ToLower(text)[0] == 'y' {
		invitedFamily = true
		fmt.Print("inviting everyone important")
		printEllipsis()
		inviteFamilyAndFriends()
	} else {
		fmt.Print("snubbing friends and relatives")
		printEllipsis()
	}

	fmt.Print("How many strangers do you want to invite: ")
	text, _ = reader.ReadString('\n')
	text = text[0 : len(text)-1]
	numStrangers, _ := strconv.ParseInt(text, 10, 32)

	if numStrangers > 0 {
		if invitedFamily {
			fmt.Println("This looks like a shindig.")
		} else {
			fmt.Println("It's going to be a rager!!!!")
		}
		time.Sleep(time.Second)

		fmt.Printf("inviting %d total strangers", numStrangers)
		printEllipsis()

		inviteStrangers(int(numStrangers))
	}

	fmt.Print("preparing the birthday party")
	printEllipsis()

	for _, c := range nonPlayableCharacters {
		somebodySays(c, "Happy Birthday Pete!")
	}
}

type character struct {
	Name       string
	Gender     string
	PosX, PosY int
	Style      tcell.Style
}

func inviteStrangers(n int) {
	for i := 0; i < n; i++ {
		n := randomName()
		g := "F"

		if rand.Int()%2 == 0 {
			g = "M"
		}

		c := randomizedCharacter(n, g)
		nonPlayableCharacters = append(nonPlayableCharacters, c)

		fmt.Printf("inviting a total stranger named %s.\n", c.Name)
		time.Sleep(100 * time.Millisecond)
	}
}

func inviteFamilyAndFriends() {
	for _, info := range []struct {
		Name   string
		Gender string
	}{
		{Name: "Erik", Gender: "M"},
		{Name: "Olivia", Gender: "F"},
		{Name: "Victoria", Gender: "F"},
		{Name: "Colin", Gender: "F"},
		{Name: "Tracy", Gender: "F"},
		{Name: "Laurie", Gender: "F"},
		{Name: "Hilary", Gender: "F"},
		{Name: "Jana", Gender: "F"},
		{Name: "Scarlet", Gender: "F"},
		{Name: "John", Gender: "M"},
		{Name: "Sheila", Gender: "F"},
		{Name: "Lulu", Gender: "F"},
		{Name: "Miles", Gender: "M"},
		{Name: "Naja", Gender: "F"},
		{Name: "Tim", Gender: "M"},
		{Name: "Elizabeth", Gender: "F"},
	} {

		fmt.Printf("inviting %s.\n", info.Name)
		time.Sleep(100 * time.Millisecond)
		nonPlayableCharacters = append(nonPlayableCharacters, randomizedCharacter(info.Name, info.Gender))
	}
}

func randomizedCharacter(name, gender string) *character {
	br := int32(50)
	r := rand.Int31()%(256-br) + br
	g := rand.Int31()%(256-br) + br
	b := rand.Int31()%(256-br) + br

	c := &character{
		Name:   name,
		Gender: gender,
		Style:  styleDefault.Foreground(tcell.NewRGBColor(r, g, b))}

	for {
		x := rand.Int() % mapWidth
		y := rand.Int() % mapHeight

		if isWalkable(x, y) {
			c.PosX = x
			c.PosY = y
			break
		}
	}

	return c
}

func (c *character) isNearby(r rune) bool {
	runes := []rune{
		getMapTile(c.PosX-1, c.PosY-1),
		getMapTile(c.PosX-1, c.PosY),
		getMapTile(c.PosX-1, c.PosY+1),
		getMapTile(c.PosX, c.PosY-1),
		getMapTile(c.PosX, c.PosY+1),
		getMapTile(c.PosX+1, c.PosY-1),
		getMapTile(c.PosX+1, c.PosY),
		getMapTile(c.PosX+1, c.PosY+1),
	}

	for _, b := range runes {
		if b == r {
			return true
		}
	}

	return false
}

func (c *character) move(x, y int) {
	if isWalkable(c.PosX+x, c.PosY+y) {
		c.PosY += y
		c.PosX += x
	}
	activateNearby(c)
}

func (c *character) detectBumps() {
	for _, c2 := range nonPlayableCharacters {
		var diffX = math.Abs(float64(c.PosX - c2.PosX))
		var diffY = math.Abs(float64(c.PosY - c2.PosY))

		if (diffX == 1 && diffY == 0) || (diffX == 0 && diffY == 1) {
			somebodyBumps(c, c2)
		}
	}
}

func characterAIHandler(doRender chan MessageDoRender, doQuit chan MessageDoQuit) {
	for {
		updateNonPlayableAI()
		doRender <- MessageDoRender{}
		time.Sleep(1776 * time.Millisecond)
	}
}

func activateNearby(c *character) {
	if c.isNearby('T') {
		alert(fmt.Sprintf("%s is eating from the table", c.Name))
	} else if c.isNearby('t') {
		alert(fmt.Sprintf("%s is using the toilet", c.Name))
	} else if c.isNearby('K') {
		alert(fmt.Sprintf("%s is working in the kitchen", c.Name))
	} else if c.isNearby('F') {
		alert(fmt.Sprintf("%s is grabbing something from the fridge", c.Name))
	} else if c.isNearby('S') {
		alert(fmt.Sprintf("%s is using a sink", c.Name))
	} else if c.isNearby('E') {
		alert(fmt.Sprintf("%s is playing the piano", c.Name))
	} else if c.isNearby('~') {
		alert(fmt.Sprintf("⌁⌁⌁ %s has had a heart attack and is being defibrillated!!!!️ ⌁⌁⌁", c.Name))
	} else if c.isNearby('H') {
		alert(fmt.Sprintf("%s is taking a shower", c.Name))
	} else if c.isNearby('C') {
		alert(fmt.Sprintf("%s is snacking at the counter", c.Name))
	} else if c.isNearby('&') {
		alert(fmt.Sprintf("%s is playing with a stuffed animal", c.Name))
	}
}

func updateNonPlayableAI() {
	for _, c := range nonPlayableCharacters {
		c.move(rand.Int()%4-2, rand.Int()%4-2)

		if rand.Int()%50 == 0 {
			somebodySays(c, randomQuote())
		}
	}
}
func playerHandleKeyEvent(key *tcell.EventKey) {
	switch key.Key() {
	case tcell.KeyUp:
		peteCharacter.move(0, -1)
	case tcell.KeyDown:
		peteCharacter.move(0, 1)
	case tcell.KeyLeft:
		peteCharacter.move(-1, 0)
	case tcell.KeyRight:
		peteCharacter.move(1, 0)
	}
	switch key.Rune() {
	case 's':
		somebodySays(peteCharacter, randomQuote())
	case 'r':
		somebodySays(peteCharacter, randomRage())
	}
	peteCharacter.detectBumps()
}

func renderCharacter(screen tcell.Screen, c *character) {
	screen.SetCell(c.PosX, c.PosY, c.Style, '@')
}

func renderCharacters(screen tcell.Screen) {
	for _, c := range nonPlayableCharacters {
		renderCharacter(screen, c)
	}
	screen.SetCell(peteCharacter.PosX, peteCharacter.PosY, peteCharacter.Style, '@')
}
