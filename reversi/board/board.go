package board

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
)

type Board struct {
  size   int
  stone  [][]stone.Stone
}

func NewBoard(size int) *Board {
  b := &Board{ size: size, stone: make([][]stone.Stone, size) }
  for i:=0; i<size; i++ {
    b.stone[i] = make([]stone.Stone, size)
    for j:=0; j<size; j++ {
      b.stone[i][j] = stone.None
    }
  }
  b.stone[b.size/2-1][b.size/2-1] = stone.White
  b.stone[b.size/2  ][b.size/2-1] = stone.Black
  b.stone[b.size/2-1][b.size/2  ] = stone.Black
  b.stone[b.size/2  ][b.size/2  ] = stone.White
  return b
}

func (b *Board) GetSize() int {
  return b.size
}

func (b *Board) GetStone(p point.Point) stone.Stone {
  if ! b.Valid(p) {
    panic("argument point is out of range")
  }
  return b.stone[p.X][p.Y]
}

func (b *Board) SetStone(p point.Point, s stone.Stone) {
  if ! b.Valid(p) {
    panic("argument point is out of range")
  }
  b.stone[p.X][p.Y] = s
}

func (b *Board) Valid(p point.Point) bool {
  return (0 <= p.X) && (p.X < b.size) && (0 <= p.Y) && (p.Y < b.size)
}


func (b *Board) Flippable(p point.Point, d point.Direction, s stone.Stone) bool {
  if q := p.Move(d); b.Valid(q) && s.Flip() == b.GetStone(q) {
    for q = q.Move(d); b.Valid(q) && stone.None != b.GetStone(q); q = q.Move(d) {
      if s == b.GetStone(q) {
        return true
      }
    }
  }
  return false
}

func (b *Board) Puttable(p point.Point, s stone.Stone) bool {
  if ! b.Valid(p) {
    return false
  }
  if b.GetStone(p) != stone.None {
    return false
  }
  var d point.Direction
  for d=0; d<point.Directions; d++ {
    if b.Flippable(p, d, s) {
      return true
    }
  }
  return false
}

func (b *Board) PutStone(p point.Point, s stone.Stone) {
  if ! b.Puttable(p, s) {
    panic("argument point is out of range")
  }
  b.SetStone(p, s)
  var d point.Direction
  for d=0; d<point.Directions; d++ {
    if b.Flippable(p, d, s) {
      for q := p.Move(d); b.Valid(q) && stone.None != b.GetStone(q) ; q = q.Move(d) {
        if s == b.GetStone(q) {
          break
        } else {
          b.SetStone(q, s) // flipping
        }
      }
    }
  }
}

func (b *Board) GetPuttablePoint(s stone.Stone) []point.Point {
  result := []point.Point{}
  for i:=0; i<b.size; i++ {
    for j:=0; j<b.size; j++ {
      p := point.Point{X:i, Y:j}
      if b.Puttable(p, s) {
        result = append(result, p)
      }
    }
  }
  return result
}

func (b *Board) CountStone(s stone.Stone) int {
  n := 0
  for i:=0; i<b.size; i++ {
    for j:=0; j<b.size; j++ {
      if b.stone[i][j] == s {
        n++
      }
    }
  }
  return n
}

func (b *Board) Copy() *Board {
  c := &Board{ size: b.size, stone: make([][]stone.Stone, b.size) }
  for i:=0; i<b.size; i++ {
    c.stone[i] = make([]stone.Stone, b.size)
    for j:=0; j<b.size; j++ {
      c.stone[i][j] = b.stone[i][j]
    }
  }
  return c
}


