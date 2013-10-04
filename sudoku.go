package main

import (
	"flag"
	"fmt"
	//	"github.com/okke-formsma/go-intset"
	"strconv"
	"strings"
)

type board struct {
	cells [81]byte
}

func (b board) String() string {
	var s = ""
	for i := 0; i < 9; i++ {
		if i > 0 && i%3 == 0 {
			s += "---+---+---\n"
		}
		for j := 0; j < 9; j++ {
			if j > 0 && j%3 == 0 {
				s += "|"
			}
			cell := b.cells[i*9+j]
			if cell == 0 {
				s += " "
			} else {
				s += fmt.Sprintf("%v", b.cells[i*9+j])
			}
		}
		s += "\n"
	}
	return s
}

func NewBoard(values string) (b board) {
	i := 0
	for _, char := range values {
		if !strings.Contains(".0123456789", string(char)) {
			continue
		}
		cell, err := strconv.ParseInt(string(char), 10, 8)
		if err == nil {
			b.cells[i] = byte(cell)
		} else {
			b.cells[i] = 0
		}
		i++
	}
	return b
}

// Returns -1 if board is full
func (b board) firstEmptyCell() int {
	for i, cell := range b.cells {
		if cell == 0 {
			return i
		}
	}
	return -1
}

func (b board) checkCell(index int) (ok bool) {
	return b.checkRow(index) && b.checkColumn(index) && b.checkSquare(index)
}

//check the row which contains index.
func (b board) checkRow(index int) (ok bool) {
	start := index - (index % 9)
	cell_value := b.cells[index]
	for i := start; i < start+9; i++ {
		if i != index && b.cells[i] == cell_value {
			return false
		}
	}
	return true
}

//check the column which contains index.
func (b board) checkColumn(index int) (ok bool) {
	cell_value := b.cells[index]
	for i := index % 9; i < 81; i += 9 {
		if i != index && b.cells[i] == cell_value {
			return false
		}
	}
	return true
}

//check the row which contains index.
func (b board) checkSquare(index int) (ok bool) {
	start_row := (index / 9) - ((index / 9) % 3)
	start_col := (index - (index % 3)) % 9

	cell_value := b.cells[index]
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			test_i := (start_row*9 + i*9) + start_col + j
			if test_i != index && b.cells[test_i] == cell_value {
				return false
			}
		}
	}
	return true

}

func (b board) Solve() (solution board, solution_found bool) {
	index := b.firstEmptyCell()
	if index == -1 {
		return b, true
	}
	var guess byte
	for guess = 1; guess < 10; guess++ {
		b.cells[index] = guess
		//fmt.Println(b)
		if b.checkCell(index) {
			solution, solution_found := b.Solve()
			if solution_found {
				return solution, true
			}
		}
	}
	// no solution found
	b.cells[index] = 0
	return b, false
}

func main() {
	example := flag.Int("example", 1, "Choose a example board between 1 and 5.")
	flag.Parse()
	examples := []string{`
	8 5 . |. . 2 |4 . . 
	7 2 . |. . . |. . 9 
	. . 4 |. . . |. . . 
	------+------+------
	. . . |1 . 7 |. . 2 
	3 . 5 |. . . |9 . . 
	. 4 . |. . . |. . . 
	------+------+------
	. . . |. 8 . |. 7 . 
	. 1 7 |. . . |. . . 
	. . . |. 3 6 |. 4 . `,
		`
	. . . |. . 6 |. . . 
	. 5 9 |. . . |. . 8 
	2 . . |. . 8 |. . . 
	------+------+------
	. 4 5 |. . . |. . . 
	. . 3 |. . . |. . . 
	. . 6 |. . 3 |. 5 4 
	------+------+------
	. . . |3 2 5 |. . 6 
	. . . |. . . |. . . 
	. . . |. . . |. . . `,
		`
	. . . |. . 5 |. 8 . 
	. . . |6 . 1 |. 4 3 
	. . . |. . . |. . . 
	------+------+------
	. 1 . |5 . . |. . . 
	. . . |1 . 6 |. . . 
	3 . . |. . . |. . 5 
	------+------+------
	5 3 . |. . . |. 6 1 
	. . . |. . . |. . 4 
	. . . |. . . |. . . `,
		"4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......",
	}

	args := flag.Args()
	var b board
	fmt.Println(args)
	if len(args) > 0 {
		fmt.Println("Using example", *example)
		b = NewBoard(args[0])
	} else if *example < len(examples) {
		b = NewBoard(examples[*example])
	}
	fmt.Println("Solving this board:")
	fmt.Println(b)
	solution, _ := b.Solve()
	fmt.Println("Solution:")
	fmt.Println(solution)
}
