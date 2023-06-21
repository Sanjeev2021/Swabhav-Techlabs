package board

import (
	"fmt"

	//board "tictactoestruct/Board"
	cell "tictactoestruct/Cell"
)

type Board struct {
	Cells [9]*cell.Cell
}

func NewBoard() *Board {
	var CellBlocks [9]*cell.Cell
	for i := 0; i < 9; i++ {
		CellBlocks[i] = cell.NewCell()
	}
	return &Board{
		Cells: CellBlocks,
	}
}

func (b *Board) ResultAnalyzer() bool {
	// check for column
	// fmt.Println("check>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	// fmt.Println(b.Cells[0].MarkedCell(), b.Cells[1].MarkedCell(), b.Cells[0].MarkedCell() == b.Cells[1].MarkedCell())
	for i := 0; i <= 8; i += 3 {
		if b.Cells[i].MarkedCell() == b.Cells[i+1].MarkedCell() && b.Cells[i+1].MarkedCell() == b.Cells[i+2].MarkedCell() && b.Cells[i].IsCellMarked() {
			return true
		}
	}
	// Checking for columns
	for i := 0; i <= 2; i++ {
		if b.Cells[i].MarkedCell() == b.Cells[i+3].MarkedCell() && b.Cells[i+3].MarkedCell() == b.Cells[i+6].MarkedCell() && b.Cells[i].IsCellMarked() {
			return true
		}
	}
	// Checking for diagonals
	if b.Cells[0].MarkedCell() == b.Cells[4].MarkedCell() && b.Cells[4].MarkedCell() == b.Cells[8].MarkedCell() && b.Cells[0].IsCellMarked() {
		return true
	}
	if b.Cells[2].MarkedCell() == b.Cells[4].MarkedCell() && b.Cells[4].MarkedCell() == b.Cells[6].MarkedCell() && b.Cells[2].IsCellMarked() {
		return true
	}
	// No winner
	return false

	//draw check

}

func (b *Board) PrintBoard() {
	fmt.Printf("\n %s | %s | %s", b.Cells[0].MarkedCell(), b.Cells[1].MarkedCell(), b.Cells[2].MarkedCell())
	fmt.Printf("\n- - - - - - - - - - - - - - -")
	fmt.Printf("\n %s | %s | %s", b.Cells[3].MarkedCell(), b.Cells[4].MarkedCell(), b.Cells[5].MarkedCell())
	fmt.Printf("\n- - - - - - - - - - - - - - - ")
	fmt.Printf("\n %s | %s | %s\n\n\n", b.Cells[6].MarkedCell(), b.Cells[7].MarkedCell(), b.Cells[8].MarkedCell())
}
