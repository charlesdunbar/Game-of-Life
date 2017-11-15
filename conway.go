// Implementation of Conway's Game of Life
// Start with ASCII, learn about graphics later

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type board [][]bool

var currentGeneration, currentPopulation, maxGeneration, maxPopulation int = 0, 0, 0, 0
var seed int64 = 1

// Print board nicely
func (b board) String() string {
	retVal := ""
	currentGeneration++
	currentPopulation = 0
	for i := range b {
		for j := range b[i] {
			if b[i][j] {
				retVal += "X "
				currentPopulation++
			} else {
				retVal += "O "
			}
		}
		retVal += "\n"
	}
	if currentPopulation > maxPopulation {
		maxPopulation = currentPopulation
		maxGeneration = currentGeneration
	}
	return retVal

}

// Deep copy function
func (b board) Copy() board {
	newBoard := make(board, len(b))
	for i := range newBoard {
		newBoard[i] = make([]bool, len(b[i]))
	}
	for i := range b {
		copy(newBoard[i], b[i])
	}
	return newBoard
}

func runRules(b *board) board {
	newBoard := b.Copy()
	for i := range *b {
		for j := range (*b)[i] {
			// Any dead cell with exactly three live neighbours
			// becomes a live cell, as if by reproduction.
			if !(*b)[i][j] {
				if numNeighbors := findNeighbors(b, i, j); numNeighbors == 3 {
					newBoard[i][j] = true
				}
			} else {
				switch numNeighbors := findNeighbors(b, i, j); numNeighbors {
				// Any live cell with fewer than two live neighbours dies,
				// as if caused by underpopulation.
				case 0, 1:
					newBoard[i][j] = false

				// Any live cell with two or three live neighbours
				// lives on to the next generation.
				case 2, 3:
					// Do nothing, just used to show the rule

				// Any live cell with more than three live neighbours
				// dies, as if by overpopulation.
				case 4, 5, 6, 7, 8:
					newBoard[i][j] = false
				}
			}
		}
	}
	return newBoard
}

// Find how many neighbors a given node has.
// Returns the number of live neighbors surrounding the node
func findNeighbors(b *board, x, y int) int {
	// Initially going to not wrap the board
	neighborCount := 0

	// To understand why (*b), read
	// https://stackoverflow.com/questions/25290956/go-update-slice-iterating-error-does-not-support-indexing

	// Starting with finding north neighbor, check clockwise for populated nodes
	// Also need to make sure we're not checking ourselves with the modified numbers

	if validLocation(b, x-1, y) {
		if (*b)[x-1][y] {
			neighborCount++
		}
	}
	if validLocation(b, x-1, y+1) {
		if (*b)[x-1][y+1] {
			neighborCount++
		}
	}
	if validLocation(b, x, y+1) {
		if (*b)[x][y+1] {
			neighborCount++
		}
	}
	if validLocation(b, x+1, y+1) {
		if (*b)[x+1][y+1] {
			neighborCount++
		}
	}
	if validLocation(b, x+1, y) {
		if (*b)[x+1][y] {
			neighborCount++
		}
	}
	if validLocation(b, x+1, y-1) {
		if (*b)[x+1][y-1] {
			neighborCount++
		}
	}
	if validLocation(b, x, y-1) {
		if (*b)[x][y-1] {
			neighborCount++
		}
	}
	if validLocation(b, x-1, y-1) {
		if (*b)[x-1][y-1] {
			neighborCount++
		}
	}

	return neighborCount
}

// Check to make sure we're not out of bounds
// Return false if x or y is outside bounds of board
func validLocation(b *board, x, y int) bool {
	if x < 0 || x >= len((*b)) || y < 0 || y >= len((*b)[x]) {
		return false
	}
	return true
}

// Make 2d board and populate each element with either false (unpopulated) or
// true (populated)
func initBoard(x, y int) board {
	retVal := make(board, x)
	for i := range retVal {
		retVal[i] = make([]bool, y)
	}
	rand.Seed(seed)
	for i := range retVal {
		for j := range retVal[i] {
			retVal[i][j] = rand.Float32() < 0.5
		}
	}
	return retVal
}

func main() {
	testBoard := initBoard(5, 5)
	for {
		fmt.Print(testBoard)
		fmt.Printf("\nCurrent Generation: %d\n", currentGeneration)
		fmt.Printf("Current Population: %d\n\n", currentPopulation)
		if currentPopulation == 0 {
			fmt.Println("Game Over")
			fmt.Printf("Seed %d lastest %d generations, with a \n", seed, currentGeneration)
			fmt.Printf("maximum population of %d during generation %d\n", maxPopulation, maxGeneration)
			os.Exit(1)
		}
		testBoard = runRules(&testBoard)
		time.Sleep(1 * time.Second)
	}

}
