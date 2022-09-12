package player

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/game"
)

type HumanPlayer struct {
}

func NewHumanPlayer() HumanPlayer {
  return HumanPlayer{}
}

func (r HumanPlayer) Play(g game.Game) (p point.Point, pass, giveup bool) {
  return point.Unavailable, false, false
}

func (r HumanPlayer) UseUIInput() bool {
  return true
}



