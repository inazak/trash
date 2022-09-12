package stone

type Stone int

const (
  None  = iota
  White
  Black
)

func (s Stone) String() string {
  switch s {
  case Black:
    return "Black"
  case White:
    return "White"
  case None:
    return "None"
  default:
    panic("stone is out of range")
  }
}

func (s Stone) Flip() Stone {
  switch s {
  case Black:
    return White
  case White:
    return Black
  default:
    panic("stone is not flippable")
  }
}


