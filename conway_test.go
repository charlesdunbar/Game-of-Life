package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNextGeneration(t *testing.T) {
	b := board{
		{false, true, true, false, true},
		{true, false, true, true, false},
		{false, false, false, false, true},
		{false, false, true, true, true},
		{true, false, true, false, false},
	}
	fmt.Println(b)

	b = runRules(&b)

	fmt.Println(b)
	// Test cell dying from not enough neighbors
	if b[4][0] {
		t.Errorf("Cell[4][0] with no neighbors should die")
	}
	// Test cell with 2 neighbors
	if !b[4][2] {
		t.Errorf("Cell[4][2] with 2 neighbors should still be alive")
	}
	// Test cell dying from too many neighbors
	if b[3][3] {
		t.Errorf("Cell[3][3] with 4 neighbors should die")
	}
	// Test dead cell becoming alive
	if !b[3][1] {
		t.Errorf("Dead cell[3][1] with 3 neighbors should live")
	}

}

func TestFindNeighbors(t *testing.T) {
	b := board{
		{false, true, true, false},
		{true, false, true, true},
		{false, false, false, false},
		{false, false, true, true},
	}

	if findNeighbors(&b, 0, 0) != 2 {
		t.Errorf("findNeighbors(b[0][0]) != 2")
	}
	if findNeighbors(&b, 0, 2) != 3 {
		t.Errorf("findNeighbors(b[0][2]) != 3")
	}
	if findNeighbors(&b, 0, 3) != 3 {
		t.Errorf("findNeighbors(b[0][3]) != 3")
	}
	if findNeighbors(&b, 2, 0) != 1 {
		t.Errorf("findNeighbors(b[2][0]) != 1")
	}
	if findNeighbors(&b, 2, 2) != 4 {
		t.Errorf("findNeighbors(b[2][2]) != 4")
	}
	if findNeighbors(&b, 3, 3) != 1 {
		t.Errorf("findNeighbors(b[3][3]) != 1")
	}
	//fmt.Println(b)
}

func TestValidLocation(t *testing.T) {

}

func TestInitBoard(t *testing.T) {
	x := rand.Intn(10)
	y := rand.Intn(10)
	testBoard := initBoard(x, y)
	if len(testBoard) != x {
		t.Errorf("len(testBoard) != %d", x)
	}
	for i := range testBoard {
		if len(testBoard[i]) != y {
			t.Errorf("len(testBoard[%d] != %d", i, y)
		}
	}

}
