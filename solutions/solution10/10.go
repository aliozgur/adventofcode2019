package solution10

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strings"
)

/*
**** Our reference  point is (x1,y1) ****

1) Pick a point (x2,y2) and find the slope (m) of the line between (x1,y1) and (x2,y2)
m = (y2 - y1) / (x2 - x1)

2) Solve the below equation (we already know m) for any of the (x1,y1) (x2,y2)
y = mx + b

3) Find all points other than (x1,y1) on the line

4) Calculate the distance of (x1,y1) to found points

5) Pick the point with minimum distance

*/

type AsteroidMap struct {
	Asteroids []*Asteroid
	MaxX      float64
	MaxY      float64
}

type Asteroid struct {
	X                 float64
	Y                 float64
	Atan2             float64
	ObservedAsteroids []*Asteroid
}

func (asteroid Asteroid) String() string {
	return fmt.Sprintf("(%f,%f)", asteroid.X, asteroid.Y)
}
func NewAsteroidMap(input string) (asteroidMap AsteroidMap) {
	asteroidMap = AsteroidMap{Asteroids: make([]*Asteroid, 0)}
	var lines = strings.NewReader(input)
	scanner := bufio.NewScanner(lines)
	y := 0.0
	x := 0.0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		x = -1
		asteroids := strings.Split(line, "")
		for _, v := range asteroids {
			x++
			if v != "#" {
				continue
			}
			asteroid := Asteroid{X: x, Y: y, Atan2:math.Atan2(y,x), ObservedAsteroids: make([]*Asteroid, 0)}
			asteroidMap.Asteroids = append(asteroidMap.Asteroids, &asteroid)

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

func (asteroid Asteroid) LineFunc(from Asteroid) (result func(a Asteroid) bool) {
	if asteroid.X == from.X {
		result = func(a Asteroid) bool {
			return asteroid.X == a.X
		}
		return
	}

	m := (from.Y - asteroid.Y) / (from.X - asteroid.X) // Slope
	b := asteroid.Y - (m * asteroid.X)

	result = func(a Asteroid) bool {
		//log.Printf("(%f,%f), (%f,%f) Line function is y = %f*x + %f\n",asteroid.X,asteroid.Y, from.X,from.Y,m,b)
		//log.Printf("Checking (%f,%f) Line function is y = %f*x + %f\n",a.X,a.Y,m,b)
		result := m*a.X + b == a.Y
		return result
	}
	return
}

func (asteroid Asteroid) IsOnLine(line func(x float64, y float64) bool) (result bool) {
	result = line(asteroid.X, asteroid.Y)
	return
}
func (asteroid Asteroid) DirectionTo(to Asteroid) (dir int){
	if asteroid.X == to.X{
		if asteroid.Y > to.Y{
			dir = 1
		} else{
			dir = -1
		}
	} else if asteroid.Y == to.Y{
		if asteroid.X > to.X{
			dir = 1
		} else{
			dir = -1
		}
	} else{
		if asteroid.X > to.X {
			dir = 1
		} else{
			dir = -1
		}
	}
	return dir
}
func (asteroidMap *AsteroidMap) Solve() (max int, maxAsteroid *Asteroid) {
	numOfAsteroids := len(asteroidMap.Asteroids)
	max = math.MinInt32
	maxAsteroid = nil

	for i := 0; i < numOfAsteroids; i++ {
		a := asteroidMap.Asteroids[i]

		for j := 0; j < numOfAsteroids; j++ {
			p1 := asteroidMap.Asteroids[j]
			if p1 == a {
				continue
			}
			dir := a.DirectionTo(*p1)
			isOnLineFunc := a.LineFunc(*p1)
			dMin := a.Distance(*p1)
			hasObstacle := false

			for k := 0; k < numOfAsteroids; k++ {
				p2 := asteroidMap.Asteroids[k]
				if p2 == a || p2 == p1 || !isOnLineFunc(*p2) || a.DirectionTo(*p2) != dir {
					continue
				}
				d := a.Distance(*p2)
				if d < dMin {
					hasObstacle = true
					break
				}
			}

			if !hasObstacle {
				a.ObservedAsteroids = append(a.ObservedAsteroids,p1)
			}
		}
		observedCnt := len(a.ObservedAsteroids)
		if observedCnt > max {
			max = observedCnt
			maxAsteroid = a
		}
	}

	return
}
