package ui

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
  "github.com/inazak/reversi/game"
)

type Controller interface {
  Init(g game.Game)
  Gameset(g game.Game, winner stone.Stone)
  Wait(g game.Game) //waiting for next move
  Pass(g game.Game)
  Giveup(g game.Game)
  Put(g game.Game, p point.Point)
  Input(g game.Game) (p point.Point, pass, giveup bool)
}

