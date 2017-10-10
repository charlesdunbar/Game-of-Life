package main

import (
	"math/rand"
	"testing"
)

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
