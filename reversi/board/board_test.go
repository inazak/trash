package board

import (
  "fmt"
  "testing"
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
)


func TestCountStone(t *testing.T) {
  b := NewBoard(8)

  if b.CountStone(stone.Black) != 2 {
    t.Errorf("black stones is 4, got unexpected")
  }
  if b.CountStone(stone.White) != 2 {
    t.Errorf("white stones is 4, got unexpected")
  }
  if b.CountStone(stone.None) != 60 {
    t.Errorf("none stones is 60, got unexpected")
  }
}

func TestPuttable(t *testing.T) {
  b := NewBoard(8)

  if b.Puttable(point.Point{X:0, Y:0}, stone.Black) {
    t.Errorf("black [0,0] is not puttable, got unexpected")
  }
  if ! b.Puttable(point.Point{X:2, Y:3}, stone.Black) {
    t.Errorf("black [2,3] is puttable, got unexpected")
  }
  if b.Puttable(point.Point{X:2, Y:3}, stone.White) {
    t.Errorf("white [2,3] is not puttable, got unexpected")
  }
  if b.Puttable(point.Point{X:3, Y:3}, stone.White) {
    t.Errorf("white [3,3] is not puttable, got unexpected")
  }
  if ! b.Puttable(point.Point{X:3, Y:5}, stone.White) {
    t.Errorf("white [3,5] is puttable, got unexpected")
  }
}

func TestGetPuttablePoint(t *testing.T) {
  b := NewBoard(8)
  ps := b.GetPuttablePoint(stone.Black)

  if len(ps) != 4 {
    t.Errorf("black has 4 puttable point, got unexpected")
  }
  if ! ps[0].Equal(point.Point{X:2, Y:3}) {
    t.Errorf("puttable point ps[0] is [2,3], got=%v", ps[0])
  }
  if ! ps[1].Equal(point.Point{X:3, Y:2}) {
    t.Errorf("puttable point ps[1] is [3,2], got=%v", ps[1])
  }
  if ! ps[2].Equal(point.Point{X:4, Y:5}) {
    t.Errorf("puttable point ps[2] is [4,5], got=%v", ps[2])
  }
  if ! ps[3].Equal(point.Point{X:5, Y:4}) {
    t.Errorf("puttable point ps[3] is [5,4], got=%v", ps[3])
  }
}


func TestPutStone(t *testing.T) {
  b := NewBoard(8)

  for i:=0; i<b.size; i++ {
    for j:=0; j<b.size; j++ {
      b.stone[i][j] = stone.None
    }
  }

  for i:=0; i<5; i++ {
    b.stone[i][0] = stone.Black
    b.stone[i][4] = stone.Black
    b.stone[0][i] = stone.Black
    b.stone[4][i] = stone.Black
  }
  for i:=1; i<4; i++ {
    b.stone[i][1] = stone.White
    b.stone[i][3] = stone.White
    b.stone[1][i] = stone.White
    b.stone[3][i] = stone.White
  }

  //debugBoardPrint(b)
  b.PutStone(point.Point{X:2, Y:2}, stone.Black)
  //debugBoardPrint(b)

  for i:=1; i<4; i++ {
    if b.stone[i][1] != stone.Black {
      t.Errorf("expected is black, but got=%v %v", b.stone[i][1], point.Point{X:i, Y:1})
    }
    if b.stone[i][3] != stone.Black {
      t.Errorf("expected is black, but got=%v %v", b.stone[i][3], point.Point{X:i, Y:3})
    }
    if b.stone[1][i] != stone.Black {
      t.Errorf("expected is black, but got=%v %v", b.stone[1][i], point.Point{X:1, Y:i})
    }
    if b.stone[3][i] != stone.Black {
      t.Errorf("expected is black, but got=%v %v", b.stone[3][i], point.Point{X:3, Y:i})
    }
  }
}

func debugBoardPrint(b *Board) {
  fmt.Printf("\n")
  for y:=0; y<b.GetSize(); y++ {
    fmt.Printf("  ")
    for x:=0; x<b.GetSize(); x++ {
      m := debugStoneToMark(b.GetStone(point.Point{X:x, Y:y}))
      fmt.Printf("%v ", m)
    }
    fmt.Printf("\n")
  }
  fmt.Printf("\n")
}

func debugStoneToMark(s stone.Stone) string {
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

