package game

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
  "github.com/inazak/reversi/board"
)

type Game8x8 struct {
  turn   stone.Stone
  board  *board.Board
}

func NewGame8x8() *Game8x8 {
  return &Game8x8{
    turn:  stone.Black,
    board: board.NewBoard(8),
  }
}

func (g *Game8x8) GetTurn() stone.Stone {
  return g.turn
}
func (g *Game8x8) FlipTurn() {
  g.turn = g.turn.Flip()
}

func (g *Game8x8) GetBoardSize() int {
  return g.board.GetSize()
}

func (g *Game8x8) GetStone(p point.Point) stone.Stone {
  return g.board.GetStone(p)
}

func (g *Game8x8) PutStone(p point.Point, s stone.Stone) {
  g.board.PutStone(p, s)
}

func (g *Game8x8) GetPuttablePoint(s stone.Stone) []point.Point {
  return g.board.GetPuttablePoint(s)
}

func (g *Game8x8) CountStone(s stone.Stone) int {
  return g.board.CountStone(s)
}

func (g *Game8x8) ValidPoint(p point.Point) bool {
  return g.board.Valid(p)
}

func (g *Game8x8) Puttable(p point.Point, s stone.Stone) bool {
  return g.board.Puttable(p, s)
}

func (g *Game8x8) Gameset() (bool, stone.Stone) {
  bp := g.board.GetPuttablePoint(stone.Black)
  wp := g.board.GetPuttablePoint(stone.White)
  if len(bp) == 0 && len(wp) == 0 {
    bc := g.board.CountStone(stone.Black)
    wc := g.board.CountStone(stone.White)
    if bc > wc {
      return true, stone.Black
    } else {
      return true, stone.White
    }
  } else {
    return false, stone.None
  }
}


func (g *Game8x8) Copy() Game {
  return &Game8x8{
    turn:  g.turn,
    board: g.board.Copy(),
  }
}


