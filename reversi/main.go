package main

import (
  "github.com/inazak/reversi/game"
  "github.com/inazak/reversi/ui"
  "github.com/inazak/reversi/player"
  "github.com/inazak/reversi/session"
)

func main() {

  game := game.NewGame6x6()
  ctrl := ui.NewCUIController()
  plyb := player.NewHumanPlayer()
  plyw := player.NewRandomPlayer()

  session := session.NewSession(game, ctrl, plyb, plyw)
  session.Start()
}

