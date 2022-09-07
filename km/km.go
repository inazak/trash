package km

import (
  "strconv"
  "strings"
  "fmt"
)

type Term interface {
  String() string
}

// de Bruijn index notation, start with 0
type Variable struct {
  Index int
}

// curring, no arguments required
type Function struct {
  Body Term
}

type Application struct {
  Left  Term
  Right Term
}

type Closure struct {
  T Term
  E []Term
}


func (v Variable) String() string {
  return strconv.Itoa(v.Index)
}

func (f Function) String() string {
  return "^." + f.Body.String()
}

func (a Application) String() string {
  return "(" + a.Left.String() + " " + a.Right.String() + ")"
}

func (c Closure) String() string {
  s := "{" + c.T.String() + ", ["
  ss := []string{}
  for i := len(c.E)-1; i >= 0; i = i-1 {
    ss = append(ss, c.E[i].String())
  }
  s += strings.Join(ss, ",") + "]}"
  return s
}


type KrivineMachine struct {
  code  Term
  stack []Term
  env   []Term
  started bool
}

func NewKrivineMachine(c Term) *KrivineMachine {
  return &KrivineMachine{ code: c }
}

type VMError struct {
  m string
}

func (e VMError) Error() string {
  return fmt.Sprintf("%s", e.m)
}

func (km *KrivineMachine) StackIsEmpty() bool {
  return len(km.stack) == 0
}

func (km *KrivineMachine) EnvIsEmpty() bool {
  return len(km.env) == 0
}

func (km *KrivineMachine) PushStack(t Term) {
  km.stack = append(km.stack, t)
}

func (km *KrivineMachine) PushEnv(t Term) {
  km.env = append(km.env, t)
}

func (km *KrivineMachine) PopStack() (t Term) {
  if km.StackIsEmpty() {
    panic(VMError{m:"PopStack fail: insufficient stack size"})
  }
  t   = km.stack[len(km.stack)-1]
  km.stack = km.stack[:len(km.stack)-1]
  return t
}

func (km *KrivineMachine) PopEnv() (t Term) {
  if km.EnvIsEmpty() {
    panic(VMError{m:"PopEnv fail: insufficient stack size"})
  }
  t   = km.env[len(km.env)-1]
  km.env = km.env[:len(km.env)-1]
  return t
}

func (km *KrivineMachine) TopStack() Term {
  if km.StackIsEmpty() {
    panic(VMError{m:"TopStack fail: insufficient stack size"})
  }
  return km.stack[len(km.stack)-1]
}

func (km *KrivineMachine) TopEnv() Term {
  if km.EnvIsEmpty() {
    panic(VMError{m:"TopEnv fail: insufficient stack size"})
  }
  return km.env[len(km.env)-1]
}

func AssertClosure(t Term) {
  if _, ok := t.(Closure) ; !ok {
    panic(VMError{m:"AssertClosure fail"})
  }
}

func (km *KrivineMachine) Done() bool {
  return km.started && km.StackIsEmpty() && km.EnvIsEmpty()
}

func (km *KrivineMachine) ReduceStep() (err error) {
  km.started = true

  // capture VMError from panic
  defer func() {
    if rec := recover(); rec != nil {
      if _, ok := rec.(VMError); ok {
        err = rec.(VMError)
      } else {
        panic(rec)
      }
    }
  }()

  switch code := km.code.(type) {

  case Application:
    km.code = code.Left
    km.PushStack(Closure{ code.Right, km.env })

  case Function:
    km.code = code.Body
    km.PushEnv(km.PopStack())

  case Variable:
    if code.Index == 0 {
      AssertClosure(km.TopEnv())
      closure, _ := km.TopEnv().(Closure)
      km.code = closure.T
      km.env  = closure.E
    } else {
      code.Index = code.Index - 1
      km.code    = code
      _ = km.PopEnv()
    }
  }

  return nil
}


func (km *KrivineMachine) Dump() string {
  d := fmt.Sprintf("Code : %v\n", km.code.String())
  d += fmt.Sprintf("Stack: ")
  for _, s := range km.stack {
    d += fmt.Sprintf("%v, ", s.String())
  }
  d += fmt.Sprintf("\nEnv  : ")
  for _, e := range km.env {
    d += fmt.Sprintf("%v, ", e.String())
  }
  return d + "\n"
}

