// 30/12/2020

package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 500
	screenPadding             = 5 // limits at screen boundaries
	stepSize                  = 1
	populationSize            = 50
	infectionRadius           = 30
	infectionRate             = 0.9                     // probability to be infected if close to an infected person
	flagVelocity              = 1                       // (1/0) used to set no velocity, if required
	timeInterval              = 2                       // interval in seconds
	resHeader                 = "Day, Infected,Healthy" // header for the output csv file
)

var (
	normalFont    font.Face
	green         = color.RGBA{10, 255, 55, 255}
	red           = color.RGBA{255, 10, 55, 255}
	population    = make([]*person, populationSize)
	grid          [screenWidth + 1][screenHeight + 1]int
	timeLastCheck time.Time
	statusColor   color.RGBA
	results       [][]string
	resultCounter int
	locker        sync.RWMutex
	wg            sync.WaitGroup
)

func init() {
	tt, err := opentype.Parse(goregular.TTF)
	if err != nil {
		fmt.Printf("Parse: %v", err)
	}
	const dpi = 72
	normalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 11,
		DPI:  dpi,
	})
	if err != nil {
		fmt.Printf("NewFace: %v", err)
	}
}

// Game struct
type Game struct{}

// Update the world state
func (g *Game) Update() error {
	for _, p := range population {
		p.move()
	}
	for _, p := range population {
		p.checkNeighbours()
	}
	if timeElapsed := time.Since(timeLastCheck); timeElapsed > timeInterval*time.Second {
		timeLastCheck = time.Now()
		var curRes []string
		curRes = append(curRes, fmt.Sprint(resultCounter))
		curRes = append(curRes, fmt.Sprint(countInfected()))
		curRes = append(curRes, fmt.Sprint(populationSize-countInfected()))
		results = append(results, curRes)
		resultCounter++
		fmt.Printf("Infected : %v\n", countInfected())
		fmt.Printf("Healthy : %v\n", populationSize-countInfected())
	}
	return nil
}

// Draw the world
func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.White)
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	for _, v := range population {
		if v.isSick {
			statusColor = red
		} else {
			statusColor = green
		}
		screen.Set(int(v.position.x)+1, int(v.position.y), statusColor)
		screen.Set(int(v.position.x)-1, int(v.position.y), statusColor)
		screen.Set(int(v.position.x), int(v.position.y)-1, statusColor)
		screen.Set(int(v.position.x), int(v.position.y)+1, statusColor)
		//text.Draw(screen, fmt.Sprint(v.id), normalFont, int(v.position.x+4), int(v.position.y+2), color.White)
	}
}

// Layout of the window
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func createGrid() {
	for i := 0; i < screenWidth+1; i++ {
		for j := 0; j < screenHeight+1; j++ {
			grid[i][j] = -1
		}
	}
}

func createPopulation() {
	var sick bool
	for i := 0; i < populationSize; i++ {
		if test := rand.Float64(); test < 0.10 {
			sick = true
		} else {
			sick = false
		}
		population[i] = createPerson(i, sick)
	}
	// fmt.Printf("Person 39 data : x=%v, y=%v, velx=%v and vely =%v\n", population[39].position.x, population[39].position.y, population[39].velocity.x, population[39].velocity.y)
}

func (p person) checkNeighbours() int {
	count := 0
	lower, upper := p.position.addV(-infectionRadius), p.position.addV(infectionRadius)
	for i := math.Max(lower.x, 0); i < math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j < math.Min(upper.y, screenHeight); j++ {
			if neighbourID := grid[int(i)][int(j)]; neighbourID != -1 && neighbourID != p.id && p.isSick {
				if dist := (population[neighbourID].position).distance(p.position); dist < infectionRadius {
					if p := population[neighbourID]; p.willGetInfection() {
						p.isSick = true
					}
					count++
				}
				//fmt.Printf("Person %v has %v neighbour(s)\n", p.id, count)
			}
		}
	}
	return count
}

func (p person) willGetInfection() bool {
	if prob := rand.Float64(); prob < infectionRate {
		return true
	}
	return false
}

func countInfected() int {
	count := 0
	for _, p := range population {
		if p.isSick {
			count++
		}
	}
	return count
}

// creates a csv file with the results
func getResults() {
	fmt.Println(len(results))
	timeStamp := fmt.Sprintf("%d-%02d-%02d",
		timeLastCheck.Year(), timeLastCheck.Month(), timeLastCheck.Day())
	filename := "test" + timeStamp + ".csv"
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println("Failed to open file", err)
	}
	w := csv.NewWriter(f)
	err = w.WriteAll(results)
	if err != nil {
		fmt.Println("Failed to write in file", err)
	}
}

func main() {
	timeLastCheck = time.Now()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Testing")
	createGrid()
	createPopulation()
	defer getResults()
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println(err)
	}
}
