package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth, screenHeight = 640, 500
	stepSize                  = 1
	populationSize            = 10
	infectionRadius           = 50
	flagVelocity              = 0 // (1/0) used to set no velocity, if required
)

var (
	normalFont font.Face
	green      = color.RGBA{10, 255, 55, 255}
	red        = color.RGBA{255, 10, 55, 255}
	population = make([]*person, populationSize)
	grid       [screenWidth + 1][screenHeight + 1]int
)

func init() {
	tt, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	const dpi = 72
	normalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 11,
		DPI:  dpi,
	})
	if err != nil {
		log.Fatalf("NewFace: %v", err)
	}
}

// Game struct
type Game struct{}

// Update the world state
func (g *Game) Update() error {
	for _, p := range population {
		p.move()
	}
	return nil
}

// Draw the world
func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.White)
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	for _, v := range population {
		statusColor := green
		if v.isSick {
			statusColor = red
		}
		screen.Set(int(v.position.x)+1, int(v.position.y), statusColor)
		screen.Set(int(v.position.x)-1, int(v.position.y), statusColor)
		screen.Set(int(v.position.x), int(v.position.y)-1, statusColor)
		screen.Set(int(v.position.x), int(v.position.y)+1, statusColor)
		text.Draw(screen, fmt.Sprint(v.id), normalFont, int(v.position.x+4), int(v.position.y+2), color.White)
	}
}

// Layout of the window
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func createGrid() {
	for i := 0; i < screenWidth; i++ {
		for j := 0; j < screenHeight; j++ {
			grid[i][j] = -1
		}
	}
}

func createPopulation() {
	var sick bool
	for i := 0; i < populationSize; i++ {
		test := rand.Float64()
		if test < 0.40 {
			sick = true
		} else {
			sick = false
		}
		population[i] = createPerson(i, sick)
	}
}

func (p person) checkNeighbours() int {
	count := 0
	lower, upper := p.position.addV(-infectionRadius), p.position.addV(infectionRadius)
	for i := math.Max(lower.x, 0); i < math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j < math.Min(upper.y, screenHeight); j++ {
			if neighbourID := grid[int(i)][int(j)]; neighbourID != -1 && neighbourID != p.id && p.isSick {
				if dist := (population[neighbourID].position).distance(p.position); dist < infectionRadius {
					count++
				}
				fmt.Printf("Person %v has %v neighbour(s)\n", p.id, count)
			}
		}
	}

	return count
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Testing")
	createGrid()
	createPopulation()
	for _, p := range population {
		p.checkNeighbours()
	}
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
