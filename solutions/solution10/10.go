package solution10

import (
	in11 "adventofcode/inputs/input11"
	sol11 "adventofcode/solutions/solution11"
	"adventofcode/utils"
	"bufio"
	"flag"
	"fmt"
	tm "github.com/buger/goterm"
	"log"
	"math"
	"sort"
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
	ObservedAsteroids map[string]Target
}

type Target struct{
	Ast *Asteroid
	Angle float64
	Distance float64
}

type ByAngleClockwise []Target

func (a ByAngleClockwise) Len() int           { return len(a) }
func (a ByAngleClockwise) Less(i, j int) bool { return a[i].Angle > a[j].Angle }
func (a ByAngleClockwise) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (target Target) String() string{
	return fmt.Sprintf("Angle:%f Distance:%f %s ", target.Angle, target.Distance,target.Ast.String())
}

func (asteroid Asteroid) String() string {
	return fmt.Sprintf("(%f,%f)", asteroid.X, asteroid.Y)
}

func NewAsteroidMap(input string, userOptions  *AsteroidMapOptions) (asteroidMap AsteroidMap) {
	var options = userOptions
	if userOptions == nil {
		boolPtr := flag.Bool("p", false, "Print on screen")
		numbPtr := flag.Int("ps", 0, "Slow down print progress")
		flag.Parse()

		options = &AsteroidMapOptions{Print: *boolPtr, PrintSleepMiliseconds: *numbPtr}
	}

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
				asteroid := Asteroid{X: x, Y: y, ObservedAsteroids: make(map[string]Target, 0)}
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

func SolveParts(){
	options := utils.UserOptions{Print:false}
	sol11.Solve(in11.Puzzle,sol11.BLACKPANEL,sol11.HEADINGUP,&options)

	options = utils.UserOptions{Print:true,PrintSleepMiliseconds:5}
	sol11.Solve(in11.Puzzle,sol11.WHITEPANEL,sol11.HEADINGUP,&options)
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
			// Change X and Y to get clockwise angles
			// Angels decrease as we move clockwise
			angle := math.Atan2(p.X-a.X,p.Y-a.Y)
			ap, ok := angles[angle]

			asteroidMap.drawPoint(*p,tm.BLUE)
			if !ok{

				angles[angle] = p
				a.ObservedAsteroids[p.String()] = Target{Angle:angle,Distance:a.Distance(*p),Ast:p}
				asteroidMap.drawPoint(*p,tm.GREEN)
			} else{
				d1 := a.Distance(*ap)
				d2 := a.Distance(*p)
				if d1 < d2 {
					angles[angle] = p
					delete(a.ObservedAsteroids,ap.String())
					a.ObservedAsteroids[p.String()] = Target{Angle:angle,Distance:a.Distance(*p),Ast:p}
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
	//fmt.Println("Number of asteroids:",len(asteroidMap.Asteroids))
	fmt.Println("Part 1 | Answer: ",max)
	//fmt.Println("Part 1 | Point",maxAsteroid)

	targets := make([]Target,0)
	for _,v := range maxAsteroid.ObservedAsteroids{
		targets = append(targets,v)
	}

	sort.Sort(ByAngleClockwise(targets))
	target := targets[199].Ast
	fmt.Println("Part 2 | Answer: ", target.X*100 + target.Y)

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
