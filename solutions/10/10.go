package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Grid struct {
	height            int
	width             int
	grid              []string
	current_position  Position
	previous_position Position
	cardinals         map[string][]int
}

type Position struct {
	X           int
	Y           int
	symbol      string
	connections []string
}

func (grid Grid) update_position(position Position) {
	grid.current_position = position
}

func NewGrid(plain string) *Grid {
	temp := strings.Split(strings.TrimRight(string(plain), "\n"), "\n")

	grid := &Grid{
		height: len(temp),
		width:  len(temp[0]),
		grid:   temp,
	}

	grid.cardinals = map[string][]int{
		"N": {-1, 0}, // N
		"E": {0, 1},  // E
		"S": {1, 0},  // S
		"W": {0, -1}, // W
	}

	grid.current_position = grid.starting_position()
	grid.current_position.symbol = grid.resolve_starting_symbol()
	grid.current_position.connections = grid.resolve_connections()

	return grid
}

func (grid Grid) starting_position() Position {
	for i := 0; i < len(grid.grid); i++ {
		for k := 0; k < len(grid.grid[i]); k++ {
			if grid.grid[i][k] == 'S' {
				return Position{
					X:      i,
					Y:      k,
					symbol: string(grid.grid[i][k]),
					// symbol: grid.resolve_starting_symbol(),
				}
			}
		}
	}

	return Position{}
}

func (grid Grid) peek(x int, y int) (Position, error) {
	x += grid.current_position.X
	y += grid.current_position.Y

	if (x > grid.width || x < 0) && (y > grid.height || y < 0) {
		return Position{}, errors.New("err")
	}

	return Position{
		X:      x,
		Y:      y,
		symbol: string(grid.grid[x][y]),
	}, nil
}

func (grid Grid) update_symbol() string {
	return string(grid.grid[grid.current_position.X][grid.current_position.Y])
}

func (grid *Grid) move(x int, y int) error {
	grid.current_position.X += x
	grid.current_position.Y += y
	grid.current_position.symbol = grid.update_symbol()
	grid.current_position.connections = grid.resolve_connections()

	if (x > grid.width || x < 0) && (y > grid.height || y < 0) {
		log.Fatal("err")
	}

	return nil
}

func (grid Grid) resolve_connections() []string {
	switch grid.current_position.symbol {
	case "L":
		return []string{"N", "E"}
	case "|":
		return []string{"N", "S"}
	case "J":
		return []string{"N", "W"}
	case "-":
		return []string{"W", "E"}
	case "F":
		return []string{"S", "E"}
	case "7":
		return []string{"S", "W"}
	default:
		return []string{}
	}
}

func (grid Grid) resolve_starting_symbol() string {
	current_position_connections := []string{}

	for cardinal_symbol, coordinates := range grid.cardinals {
		advanced_position, err := grid.peek(coordinates[0], coordinates[1])
		// fmt.Fprintf(os.Stderr, "%+v | %+v | %+v | (%+v,%+v)=%+v | (%+v,%+v) \n", cardinal_symbol, advanced_position.symbol, coordinates, advanced_position.X, advanced_position.Y, string(grid.grid[advanced_position.X][advanced_position.Y]), grid.current_position.X, grid.current_position.Y)
		if err != nil {
			continue
		}

		switch cardinal_symbol {
		case "N":
			// needs to have a open south, since we are north position of the starting one
			// . . . . x . { F, 7, | }
			// . x . . p .
			// . . . . . .
			if advanced_position.symbol == "F" || advanced_position.symbol == "7" || advanced_position.symbol == "|" {
				current_position_connections = append(current_position_connections, "N")
			}
		case "E":
			if advanced_position.symbol == "J" || advanced_position.symbol == "7" || advanced_position.symbol == "-" {
				current_position_connections = append(current_position_connections, "E")
			}
		case "S":
			if advanced_position.symbol == "L" || advanced_position.symbol == "J" || advanced_position.symbol == "|" {
				current_position_connections = append(current_position_connections, "S")
			}
		case "W":
			if advanced_position.symbol == "L" || advanced_position.symbol == "F" || advanced_position.symbol == "-" {
				current_position_connections = append(current_position_connections, "W")
			}
		default:
			log.Fatal("DEBUGPRINT[7]: 10.go:137: unreachable")
		}
	}

	str := strings.Join(current_position_connections, "")
	N := strings.Index(str, "N")
	E := strings.Index(str, "E")
	S := strings.Index(str, "S")
	W := strings.Index(str, "W")

	if N != -1 && E != -1 {
		return "L"
	}
	if N != -1 && S != -1 {
		return "|"
	}
	if N != -1 && W != -1 {
		return "J"
	}
	if E != -1 && W != -1 {
		return "-"
	}
	if E != -1 && S != -1 {
		return "F"
	}
	if S != -1 && W != -1 {
		return "7"
	}
	return "."
}

func (grid *Grid) move_north() {
	grid.move(grid.cardinals["N"][0], grid.cardinals["N"][1])
}
func (grid *Grid) move_east() {
	grid.move(grid.cardinals["E"][0], grid.cardinals["E"][1])
}
func (grid *Grid) move_south() {
	grid.move(grid.cardinals["S"][0], grid.cardinals["S"][1])
}
func (grid *Grid) move_west() {
	grid.move(grid.cardinals["W"][0], grid.cardinals["W"][1])
}

func (grid *Grid) traverse() {
	count := 0

	switch grid.current_position.connections[0] {
	case "N":
		grid.move_north()
	case "E":
		grid.move_east()
	case "S":
		grid.move_south()
	case "W":
		grid.move_west()
	default:
		log.Fatal("unreachable")
	}

  fmt.Fprintf(os.Stderr, "DEBUGPRINT[10]: 10.go:207: count=%+v\n", count)
}

func main() {
	temp, err := os.ReadFile("inputs/10.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := NewGrid(string(temp))
  fmt.Fprintf(os.Stderr, "DEBUGPRINT[12]: 10.go:225: grid.current_position=%+v\n", grid.current_position)
  grid.traverse()
  fmt.Fprintf(os.Stderr, "DEBUGPRINT[12]: 10.go:225: grid.current_position=%+v\n", grid.current_position)
}
