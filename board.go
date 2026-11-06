package main

import "fmt"

type Color int

const (
	White Color = iota
	Black
	None
)

type PieceType int

const (
	Empty PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Type  PieceType
	Color Color
}

type Square struct {
	x int
	y int
}

type Board [8][8]Piece

type GameState struct {
	Board          Board
	Turn           Color
	FullMoveNumber int
}

func NewGame() *GameState {
	gs := &GameState{
		Turn:           White,
		FullMoveNumber: 1,
	}

	backRow := []PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

	for i := 0; i < 8; i++ {
		gs.Board[0][i] = Piece{Type: backRow[i], Color: White}
		gs.Board[1][i] = Piece{Type: Pawn, Color: White}
		gs.Board[2][i] = Piece{Type: Empty, Color: None}
		gs.Board[3][i] = Piece{Type: Empty, Color: None}
		gs.Board[4][i] = Piece{Type: Empty, Color: None}
		gs.Board[5][i] = Piece{Type: Empty, Color: None}
		gs.Board[6][i] = Piece{Type: Pawn, Color: Black}
		gs.Board[7][i] = Piece{Type: backRow[i], Color: Black}
	}
	return gs
}

func (gs *GameState) PrintBoard() {
	fmt.Println("Board")
	symbols := map[Piece]string{
		{Pawn, White}: "♟", {Knight, White}: "♞", {Bishop, White}: "♝",
		{Rook, White}: "♜", {Queen, White}: "♛", {King, White}: "♚",
		{Pawn, Black}: "♙", {Knight, Black}: "♘", {Bishop, Black}: "♗",
		{Rook, Black}: "♖", {Queen, Black}: "♕", {King, Black}: "♔",
		{Empty, None}: "_",
	}

	fmt.Println("_ a b c d e f g h")
	for row := 0; row < 8; row++ {
		fmt.Printf("%d ", row+1)
		for column := 0; column < 8; column++ {
			fmt.Printf("%s ", symbols[gs.Board[row][column]])
			if column == 7 {
				fmt.Print("\n")
			}
		}
	}

	fmt.Println("_ a b c d e f g h")
}
