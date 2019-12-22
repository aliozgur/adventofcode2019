package solution10

import (
	"bufio"
	"fmt"
	tm "github.com/buger/goterm"
	"log"
	"math"
	"strings"
	"time"
)

type AsteroidMapOptions struct {
	Print bool
	PrintSleepMiliseconds int
}

type AsteroidMap struct {
	Asteroids []*Asteroid
	MaxX      float64
	MaxY      float64
	Options   *AsteroidMapOptions
}

type Asteroid struct {
	X                 float64
	Y                 float64
	ObservedAsteroids map[string]*Asteroid
}

func (asteroid Asteroid) String() string {
	return fmt.Sprintf("(%f,%f)", asteroid.X, asteroid.Y)
}

func NewAsteroidMap(input string, options  *AsteroidMapOptions) (asteroidMap AsteroidMap) {
	asteroidMap = AsteroidMap{Asteroids: make([]*Asteroid, 0),Options:options}
	var lines = strings.NewReader(input)
	scanner := bufio.NewScanner(lines)
	y := 0.0
	x := 0.0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		x = 0
		asteroids := strings.Split(line, "")
		for _, v := range asteroids {
			if v == "#" {
				asteroid := Asteroid{X: x, Y: y, ObservedAsteroids: make(map[string]*Asteroid, 0)}
				asteroidMap.Asteroids = append(asteroidMap.Asteroids, &asteroid)
			}
			x++
		}
		y++
	}

	asteroidMap.MaxY = y
	asteroidMap.MaxX = x
	if err := scanner.Err(); err != nil {
		log.Fatal("Can not scan inputs!")
	}
	return
}

func (asteroid Asteroid) Distance(from Asteroid) (result float64) {

	colDifPow := math.Pow(from.X-asteroid.X, 2.0)
	rowDifPow := math.Pow(from.Y-asteroid.Y, 2.0)

	result = math.Sqrt(colDifPow + rowDifPow)
	return
}

func (asteroidMap *AsteroidMap) Solve() (max int, maxAsteroid *Asteroid) {
	numOfAsteroids := len(asteroidMap.Asteroids)
	max = math.MinInt32

	for i := 0; i < numOfAsteroids; i++{
		a := asteroidMap.Asteroids[i]
		angles := make(map[float64]*Asteroid)
		asteroidMap.printMap(nil)

		asteroidMap.drawPoint(*a,tm.RED)

		for j := 0; j < numOfAsteroids; j++ {
			p := asteroidMap.Asteroids[j]
			if a == p {
				continue
			}

			angle := math.Mod(math.Atan2(-1*(p.Y-a.Y), p.X-a.X),2 * math.Pi)
			ap, ok := angles[angle]

			asteroidMap.drawPoint(*p,tm.BLUE)
			if !ok{
				angles[angle] = p
				a.ObservedAsteroids[p.String()] = p
				asteroidMap.drawPoint(*p,tm.GREEN)
			} else{
				d1 := a.Distance(*ap)
				d2 := a.Distance(*p)
				if d1 < d2 {
					angles[angle] = p
					delete(a.ObservedAsteroids,ap.String())
					a.ObservedAsteroids[p.String()] = p
					asteroidMap.drawPoint(*ap,tm.WHITE)
					asteroidMap.drawPoint(*p,tm.GREEN)
				}
			}
		}

		observedCnt := len(angles)
		if observedCnt > max {
			max = observedCnt
			maxAsteroid = a
		}

		if asteroidMap.Options != nil && asteroidMap.Options.Print {
			fmt.Println("")
			fmt.Println("Current Max", max)
			fmt.Println("Current max point", maxAsteroid)
			time.Sleep(100*time.Millisecond)
		}
	}

	asteroidMap.printMap(maxAsteroid)
	fmt.Println("")
	fmt.Println("Number of asteroids:",len(asteroidMap.Asteroids))
	fmt.Println("Final Max",max)
	fmt.Println("Final max point",maxAsteroid)
	return
}


func  (asteroidMap *AsteroidMap) drawPoint(p Asteroid, color int){
	if asteroidMap.Options == nil || !asteroidMap.Options.Print{
		return
	}

	tm.MoveCursor(int(p.X+1),int(p.Y+1))
	tm.Flush()
	tm.Print(tm.Color("●", color))
	tm.Flush()
	time.Sleep(time.Duration(asteroidMap.Options.PrintSleepMiliseconds)*time.Millisecond)
}

func (asteroidMap *AsteroidMap) printMap( a *Asteroid){
	if !asteroidMap.Options.Print{
		return
	}
	tm.Clear()
	tm.MoveCursor(1,1)
	tm.Flush()

	for x := 1; x <= int(asteroidMap.MaxX); x++{
		for y := 1; y <= int(asteroidMap.MaxY); y++{
			tm.MoveCursor(x,y)
			tm.Print(tm.Color(".",tm.WHITE))
			tm.Flush()
		}
	}
	tm.MoveCursor(1,1)
	tm.Flush()
	for _,p := range asteroidMap.Asteroids{
		tm.MoveCursor(int(p.X+1),int(p.Y+1))
		if a != nil {
			if p == a {
				tm.Print(tm.Color("●", tm.RED))
			} else if _, ok := a.ObservedAsteroids[p.String()];ok {
				tm.Print(tm.Color("●", tm.GREEN))
			} else {
				tm.Print(tm.Color("●", tm.WHITE))
			}
		} else {
			tm.Print("●")
		}
		tm.Flush()
	}
}
