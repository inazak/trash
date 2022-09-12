package player

import (
  "time"
  "math/rand"
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/game"
)

func init() {
  rand.Seed(time.Now().UnixNano())
}

type RandomPlayer struct {
}

func NewRandomPlayer() RandomPlayer {
  return RandomPlayer{}
}

func (r RandomPlayer) Play(g game.Game) (p point.Point, pass, giveup bool) {
  time.Sleep(time.Second * 1)
  ps := g.GetPuttablePoint(g.GetTurn())

  if len(ps) == 0 {
    return point.Point{ X:-1, Y:-1}, true, false //pass
  }

  p   = ps[ rand.Intn(len(ps)) ]
  return p, false, false
}

func (r RandomPlayer) UseUIInput() bool {
  return false
}

