package ui

import (
  "fmt"
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
  "github.com/inazak/reversi/game"
)

type CUIController struct {}

func NewCUIController() CUIController {
  return CUIController{}
}

func (c CUIController) Init(g game.Game) {
  printInfo(g)
}

func (c CUIController) Gameset(g game.Game, winner stone.Stone) {
  fmt.Printf("Gameset winner is %v\n", winner)
  fmt.Printf("\n")
}

func (c CUIController) Wait(g game.Game)  {
  fmt.Printf("waiting %v move ...\n", g.GetTurn())
  fmt.Printf("\n")
}

func (c CUIController) Pass(g game.Game) {
  fmt.Printf("pass %v ...\n", g.GetTurn())
  fmt.Printf("\n")
}

func (c CUIController) Giveup(g game.Game) {
  fmt.Printf("giveup %v ...\n", g.GetTurn())
  fmt.Printf("\n")
}

func (c CUIController) Put(g game.Game, p point.Point) {
  printInfo(g)
  s, _, _ := pointToFormat1A(p)
  fmt.Printf("%v move %v\n", g.GetTurn(), s)
  fmt.Printf("\n")
}

func (c CUIController) Input(g game.Game) (p point.Point, pass, giveup bool) {

  for {
    fmt.Printf("input move [1..][a..] as 1a,2b,... or `pass` or `giveup`: ")
    s := ""
    fmt.Scanf("%s\n", &s)

    if s == "pass" {
      return point.Unavailable, true, false
    }
    if s == "giveup" {
      return point.Unavailable, false, true
    }

    p, err := format1AtoPoint(s);
    if err != nil {
      fmt.Printf("error: %v\n\n", err)
      continue
    } else if ! g.Puttable(p, g.GetTurn()) {
      s, _, _ := pointToFormat1A(p)
      fmt.Printf("error: %v is not puttable\n\n", s)
      continue
    } else {
      return p, false, false
    }
  }
}

func printInfo(g game.Game) {
  fmt.Printf("\n")
  fmt.Printf("Turn: %v\n", g.GetTurn())
  fmt.Printf("\n")

  // draw board
  fmt.Printf("    ")
  for w:=0; w<g.GetBoardSize(); w++ {
    _, xs, _ := pointToFormat1A(point.Point{ X:w, Y:-1 })
    fmt.Printf("%v ", xs)
  }
  fmt.Printf("\n")

  for y:=0; y<g.GetBoardSize(); y++ {
    _, _, ys := pointToFormat1A(point.Point{ X:-1, Y:y })
    fmt.Printf(" %v  ", ys)
    for x:=0; x<g.GetBoardSize(); x++ {
      m := stoneToMark(g.GetStone(point.Point{X:x, Y:y}))
      fmt.Printf("%v ", m)
    }
    fmt.Printf("\n")
  }
  fmt.Printf("\n")
}

func stoneToMark(s stone.Stone) string {
  switch s {
  case stone.None:
    return "_"
  case stone.Black:
    return "x"
  case stone.White:
    return "o"
  default:
    panic("stone type is out of range")
  }
}


// "1a" => Point{ X: 0, Y: 0 }
// "2b" => Point{ X: 1, Y: 1 }
func pointToFormat1A(p point.Point) (s, xs, ys string) {
  xs = string(byte(p.X + 0x31))
  ys = string(byte(p.Y + 0x61))
  s  = xs + ys
  return s, xs, ys
}

// Point{ X: 0, Y: 0 } => "1a"
// Point{ X: 1, Y: 1 } => "2b"
func format1AtoPoint(s string) (point.Point, error) {
  p := point.Unavailable
  if len(s) != 2 {
    return p, fmt.Errorf("unexpected format")
  }
  if s[0] < '0' || '8' < s[0] { //LIMIT 8x8
    return p, fmt.Errorf("unexpected format")
  }
  if s[1] < 'a' || 'h' < s[1] { //LIMIT 8x8
    return p, fmt.Errorf("unexpected format")
  }

  p.X = int(s[0] - 0x31) // '1' => 0
  p.Y = int(s[1] - 0x61) // 'a' => 0

  return p, nil
}


