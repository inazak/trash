package player

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/game"
)

type Player interface {
  Play(g game.Game) (p point.Point, pass, giveup bool)
  UseUIInput() bool
}

