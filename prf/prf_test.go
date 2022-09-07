package prf

import (
  "testing"
)


func TestAdd(t *testing.T) {

  debug = false

  ps := [][]int{
    {0, 0, 0},
    {1, 0, 1},
    {0, 1, 1},
    {1, 1, 2},
    {1, 2, 3},
    {2, 1, 3},
    {2, 2, 4},
    {2, 3, 5},
    {3, 2, 5},
    {3, 3, 6},
  }

  for _, p := range ps {
    x, y, expect := p[0], p[1], p[2]
    if Add(x,y) != expect {
      t.Errorf("expect %v+%v=%v, but got=%v", x, y, expect, Add(x,y))
    }
  }

}


func TestDecrement(t *testing.T) {

  debug = false

  ps := [][]int{
    {0, 0},
    {1, 0},
    {2, 1},
    {3, 2},
  }

  for _, p := range ps {
    x, expect := p[0], p[1]
    if Decrement(x) != expect {
      t.Errorf("expect decrement(%v)=%v, but got=%v", x, expect, Decrement(x))
    }
  }

}


func TestMultiply(t *testing.T) {

  debug = false

  ps := [][]int{
    {0, 0, 0},
    {1, 0, 0},
    {0, 1, 0},
    {1, 1, 1},
    {1, 2, 2},
    {2, 1, 2},
    {2, 2, 4},
    {2, 3, 6},
    {3, 2, 6},
    {3, 3, 9},
  }

  for _, p := range ps {
    x, y, expect := p[0], p[1], p[2]
    if Multiply(x,y) != expect {
      t.Errorf("expect %v*%v=%v, but got=%v", x, y, expect, Multiply(x,y))
    }
  }

}


func TestSubtract(t *testing.T) {

  debug = false

  ps := [][]int{
    {0, 0, 0},
    {1, 0, 1},
    {0, 1, 0},
    {1, 1, 0},
    {1, 2, 0},
    {2, 1, 1},
    {2, 2, 0},
    {2, 3, 0},
    {3, 2, 1},
    {4, 2, 2},
    {5, 2, 3},
  }

  for _, p := range ps {
    x, y, expect := p[0], p[1], p[2]
    if Subtract(x,y) != expect {
      t.Errorf("expect %v-%v=%v, but got=%v", x, y, expect, Subtract(x,y))
    }
  }

}


func TestDivide(t *testing.T) {

  debug = false

  ps := [][]int{
    {0, 1, 0},
    {0, 2, 0},
    {1, 1, 1},
    {1, 2, 0},
    {2, 1, 2},
    {1, 3, 0},
    {3, 1, 3},
    {2, 2, 1},
    {4, 2, 2},
    {6, 2, 3},
    {8, 2, 4},
    {9, 3, 3},
  }

  for _, p := range ps {
    x, y, expect := p[0], p[1], p[2]
    if Divide(x,y) != expect {
      t.Errorf("expect %v/%v=%v, but got=%v", x, y, expect, Divide(x,y))
    }
  }

}


