package game

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
)

type Game interface {
  GetTurn() stone.Stone
  FlipTurn()
  GetBoardSize() int
  GetStone(p point.Point) stone.Stone
  PutStone(p point.Point, s stone.Stone)
  GetPuttablePoint(s stone.Stone) []point.Point
  CountStone(s stone.Stone) int
  ValidPoint(p point.Point) bool
  Puttable(p point.Point, s stone.Stone) bool
  Gameset() (bool, stone.Stone) //winner
  Copy() Game
}


