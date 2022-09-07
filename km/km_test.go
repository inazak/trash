package km

import (
  "testing"
  "fmt"
)

func TestReduceStep1(t *testing.T) {

  // (^.0 ^.0) => ^0
  expr := Application{ Left: Function{ Body: Variable { Index: 0 }, },
                       Right: Function{ Body: Variable { Index: 0 }, } }

  km := NewKrivineMachine( expr )
  fmt.Printf("%v-----\n", km.Dump())

  for ! km.Done() {
    err := km.ReduceStep()
    if err != nil {
      t.Errorf("TestReduceStep1: %v", err)
    }
    fmt.Printf("%v-----\n", km.Dump())
  }

}

func TestReduceStep2(t *testing.T) {

  // ((^.^.(0 1) ^.0) ^.0) => ^.0
  expr := Application{ Left:
          Application{ Left:
          Function { Body:
          Function { Body:
          Application{ Left: Variable { Index: 0 }, Right: Variable { Index: 1 } } } },
          Right: Function { Body: Variable { Index: 0 } }, },
          Right: Function { Body: Variable { Index: 0 } }, }

  km := NewKrivineMachine( expr )
  fmt.Printf("%v-----\n", km.Dump())

  for ! km.Done() {
    err := km.ReduceStep()
    if err != nil {
      t.Errorf("TestReduceStep2: %v", err)
    }
    fmt.Printf("%v-----\n", km.Dump())
  }

}

