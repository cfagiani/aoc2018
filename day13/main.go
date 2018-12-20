package main

import (
	"github.com/cfagiani/aoc2018/util"
	"strings"
	"sort"
	"fmt"
)

type Cart struct {
	PosX              int
	PosY              int
	Dir               byte
	IntersectionCount int
}

type CartList []*Cart

func main() {
	inputString := util.ReadFileAsString("input/day13.input")
	trackMap, carts := initialize(inputString)
	part1(trackMap, carts)
	trackMap, carts = initialize(inputString)
	part2(trackMap, carts)
}

func part1(trackMap [][]byte, carts []*Cart) {
	hasCollision := false
	x := -1
	y := -1

	for ; !hasCollision; hasCollision, x, y = checkCollision(carts) {
		carts = tick(trackMap, carts, false)
	}
	checkCollision(carts)
	fmt.Printf("The first collision occured at %d,%d\n", x, y)
}

func part2(trackMap [][]byte, carts []*Cart) {
	for len(carts) > 1 {
		carts = tick(trackMap, carts, true)
	}
	fmt.Printf("Last cart ended up at %d,%d\n", carts[0].PosX, carts[0].PosY)
}

func checkCollision(carts []*Cart) (bool, int, int) {
	sort.Sort(CartList(carts))
	for i := 0; i < len(carts)-1; i++ {
		if carts[i].isCollided(*carts[i+1]) {
			return true, carts[i].PosX, carts[i].PosY
		}
	}
	return false, -1, -1
}

func tick(trackMap [][]byte, carts []*Cart, removeCollisions bool) []*Cart {
	sort.Sort(CartList(carts))
	var nonRemovedCarts []*Cart
	for i := 0; i < len(carts); i++ {
		nonRemovedCarts = append(nonRemovedCarts, carts[i])
	}
	for i := 0; i < len(carts); i++ {
		advanceCart(carts[i], trackMap)
		if removeCollisions {
			nonRemovedCarts = removeCollided(nonRemovedCarts)
		}
	}
	return nonRemovedCarts
}

func advanceCart(cart *Cart, track [][]byte) {
	curPos := track[cart.PosY][cart.PosX]
	if curPos == '/' {
		if cart.Dir == '^' || cart.Dir == 'v' {
			cart.turnRight()
		} else if cart.Dir == '>' || cart.Dir == '<' {
			cart.turnLeft()
		}
	} else if curPos == '\\' {
		if cart.Dir == '^' || cart.Dir == 'v' {
			cart.turnLeft()
		} else if cart.Dir == '>' || cart.Dir == '<' {
			cart.turnRight()
		}
	} else if curPos == '+' {
		if cart.IntersectionCount%3 == 0 {
			cart.turnLeft()
		} else if cart.IntersectionCount%3 == 2 {
			cart.turnRight()
		}
		cart.IntersectionCount++
	}
	cart.advance()
}

func initialize(input string) ([][]byte, []*Cart) {
	lines := strings.Split(input, "\n")
	var carts []*Cart
	trackOnly := make([][]byte, len(lines))
	for i := 0; i < len(lines); i++ {
		trackOnly[i] = make([]byte, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			if isCart(lines[i][j]) {
				cart := NewCart(i, j, lines[i][j])
				carts = append(carts, &cart)
				trackOnly[i][j] = cart.getTrack()
			} else {
				trackOnly[i][j] = lines[i][j]
			}
		}
	}
	return trackOnly, carts
}

func isCart(r byte) bool {
	return '^' == r || '>' == r || '<' == r || 'v' == r
}

func removeCollided(carts []*Cart) []*Cart {
	var result []*Cart
	for i := 0; i < len(carts); i++ {
		hasCollision := false
		for j := 0; j < len(carts); j++ {
			if carts[i] != carts[j] && carts[i].isCollided(*carts[j]) {
				hasCollision = true
			}
		}
		if !hasCollision {
			result = append(result, carts[i])
		}
	}
	return result
}

func NewCart(i int, j int, r byte) Cart {
	return Cart{Dir: r, PosX: j, PosY: i}
}

func (c Cart) getTrack() byte {
	if c.Dir == '^' || c.Dir == 'v' {
		return '|'
	} else {
		return '-'
	}
}

func (c *Cart) turnLeft() {
	if c.Dir == '^' {
		c.Dir = '<'
	} else if c.Dir == 'v' {
		c.Dir = '>'
	} else if c.Dir == '>' {
		c.Dir = '^'
	} else {
		c.Dir = 'v'
	}
}

func (c *Cart) turnRight() {
	if c.Dir == '^' {
		c.Dir = '>'
	} else if c.Dir == 'v' {
		c.Dir = '<'
	} else if c.Dir == '>' {
		c.Dir = 'v'
	} else {
		c.Dir = '^'
	}
}

func (c *Cart) advance() {
	if c.Dir == '^' {
		c.PosY--
	} else if c.Dir == 'v' {
		c.PosY++
	} else if c.Dir == '>' {
		c.PosX++
	} else {
		c.PosX--
	}
}

func (c Cart) isCollided(other Cart) bool {
	return c.PosY == other.PosY && c.PosX == other.PosX
}

func (c Cart) toString() string {
	return fmt.Sprintf("%c %d,%d", c.Dir, c.PosX, c.PosY)
}

func (a CartList) Len() int { return len(a) }
func (a CartList) Less(i, j int) bool {
	if a[i].PosY == a[j].PosY {
		return a[i].PosX < a[j].PosX
	} else {
		return a[i].PosY < a[j].PosY
	}
}
func (a CartList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
