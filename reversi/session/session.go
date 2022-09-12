package session

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
  "github.com/inazak/reversi/game"
  "github.com/inazak/reversi/ui"
  "github.com/inazak/reversi/player"
)

type Session struct {
  game  game.Game
  ctrl  ui.Controller
  plyb  player.Player
  plyw  player.Player
}

func NewSession(g game.Game, c ui.Controller, pb, pw player.Player) Session {
  return Session { game: g, ctrl: c, plyb: pb, plyw: pw }
}

func (s Session) Start() {

  s.ctrl.Init(s.game.Copy())

  for {

    if ok, winner := s.game.Gameset(); ok {
      s.ctrl.Gameset(s.game.Copy(), winner)
      break
    }

    s.ctrl.Wait(s.game.Copy())

    var p point.Point
    var pass, giveup bool

    if s.game.GetTurn() == stone.Black {
      if s.plyb.UseUIInput() {
        p, pass, giveup = s.ctrl.Input(s.game.Copy())
      } else {
        p, pass, giveup = s.plyb.Play(s.game.Copy())
      }
    } else {
      if s.plyw.UseUIInput() {
        p, pass, giveup = s.ctrl.Input(s.game.Copy())
      } else {
        p, pass, giveup = s.plyw.Play(s.game.Copy())
      }
    }

    if pass {
      s.ctrl.Pass(s.game.Copy())
      s.game.FlipTurn() // turn change
      continue
    }

    if giveup {
      s.ctrl.Giveup(s.game.Copy())
      break
    }

    if ! s.game.Puttable(p, s.game.GetTurn()) {
      panic("player return unputtable point " + p.String())
    }

    s.game.PutStone(p, s.game.GetTurn())
    s.ctrl.Put(s.game.Copy(), p)

    s.game.FlipTurn() // turn change

  }
}

