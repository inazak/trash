package game

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
  "github.com/inazak/reversi/board"
)

type Game6x6 struct {
  turn   stone.Stone
  board  *board.Board
}

func NewGame6x6() *Game6x6 {
  return &Game6x6{
    turn:  stone.Black,
    board: board.NewBoard(6),
  }
}

func (g *Game6x6) GetTurn() stone.Stone {
  return g.turn
}
func (g *Game6x6) FlipTurn() {
  g.turn = g.turn.Flip()
}

func (g *Game6x6) GetBoardSize() int {
  return g.board.GetSize()
}

func (g *Game6x6) GetStone(p point.Point) stone.Stone {
  return g.board.GetStone(p)
}

func (g *Game6x6) PutStone(p point.Point, s stone.Stone) {
  g.board.PutStone(p, s)
}

func (g *Game6x6) GetPuttablePoint(s stone.Stone) []point.Point {
  return g.board.GetPuttablePoint(s)
}

func (g *Game6x6) CountStone(s stone.Stone) int {
  return g.board.CountStone(s)
}

func (g *Game6x6) ValidPoint(p point.Point) bool {
  return g.board.Valid(p)
}

func (g *Game6x6) Puttable(p point.Point, s stone.Stone) bool {
  return g.board.Puttable(p, s)
}

func (g *Game6x6) Gameset() (bool, stone.Stone) {
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


func (g *Game6x6) Copy() Game {
  return &Game6x6{
    turn:  g.turn,
    board: g.board.Copy(),
  }
}


