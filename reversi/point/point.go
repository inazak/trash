package point

import (
  "fmt"
)

type Point struct {
  X int
  Y int
}

type Direction int

const (
  North     = iota
  NorthEast
  East
  SouthEast
  South
  SouthWest
  West
  NorthWest
  Directions
)

var Unavailable = Point { X: -1, Y: -1 }

func (p Point) String() string {
  return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}

func (p Point) Equal(q Point) bool {
  return (p.X == q.X) && (p.Y == q.Y)
}

func (p Point) Move(d Direction) Point {
  switch d {
  case North:
    return Point{ X: p.X  , Y: p.Y-1 }
  case NorthEast:
    return Point{ X: p.X+1, Y: p.Y-1 }
  case East:
    return Point{ X: p.X+1, Y: p.Y   }
  case SouthEast:
    return Point{ X: p.X+1, Y: p.Y+1 }
  case South:
    return Point{ X: p.X  , Y: p.Y+1 }
  case SouthWest :
    return Point{ X: p.X-1, Y: p.Y+1 }
  case West:
    return Point{ X: p.X-1, Y: p.Y   }
  case NorthWest:
    return Point{ X: p.X-1, Y: p.Y-1 }
  default:
    panic("argument direction is out of range")
  }
}


