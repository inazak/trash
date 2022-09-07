package prf

import (
  "fmt"
)

type F func(v ...int) int

func value_split(v []int) ([]int, int) {
  w := make([]int, len(v))
  copy(w, v)

  list_without_last := w[:len(w)-1]
  last              := w[len(w)-1]

  return list_without_last, last
}

var debug bool

func watch(s string, v ...interface{}) {
  if debug {
    fmt.Printf(fmt.Sprintf("[watch] %s", s), v...)
  }
}


// ----- prf -----

func zero(v ...int) int {
  watch("call zero\n")
  return 0
}

func increment(n int) int {
  watch("call increment (%v)\n", n)
  return n + 1
}

func recurse(g F, h F, v ...int) int {
  list_without_last, last := value_split(v)
  watch("call recurse (%v)\n", append(list_without_last, last))

  if last == 0 {

    watch("call func g (%v)\n", list_without_last)
    return g(list_without_last...)

  } else {
    arg := append(list_without_last, last -1)
    r   := recurse(g, h, arg...)

    watch("recurce returned value=%v\n", r)
    watch("call func h (%v)\n", append(arg, r))

    return h(append(arg, r)...)
  }
}

func Add_G(v ...int) int {
  _, last := value_split(v)
  return last
}

func Add_H(v ...int) int {
  _, last := value_split(v)
  return increment(last)
}

func Add(x, y int) int {
  watch("-> call Add (%v,%v)\n", x, y)
  r := recurse(Add_G, Add_H, x, y)
  watch("=> Add returned value=%v\n", r)
  return r
}


func Decrement_G(v ...int) int {
  return zero(v...)
}

func Decrement_H(v ...int) int {
  list, _ := value_split(v)
  return list[len(list)-1]
}

func Decrement(x int) int {
  watch("-> call Decrement (%v)\n", x)
  r := recurse(Decrement_G, Decrement_H, x)
  watch("=> Decrement returned value=%v\n", r)
  return r
}


func Multiply_G(v ...int) int {
  return zero(v...)
}

func Multiply_H(v ...int) int {
  list, last := value_split(v)
  return Add(list[len(list)-2], last)
}

func Multiply(x, y int) int {
  watch("-> call Multiply (%v,%v)\n", x, y)
  r := recurse(Multiply_G, Multiply_H, x, y)
  watch("=> Multiply returned value=%v\n", r)
  return r
}


func Subtract_G(v ...int) int {
  _, last := value_split(v)
  return last
}

func Subtract_H(v ...int) int {
  _, last := value_split(v)
  return Decrement(last)
}

func Subtract(x, y int) int {
  watch("-> call Subtract (%v,%v)\n", x, y)
  r := recurse(Subtract_G, Subtract_H, x, y)
  watch("=> Subtract returned value=%v\n", r)
  return r
}



// ----- minimize -----

func minimize(f func(int) int) int {
  n := 0
  for {
    watch("-> call minimize function n=%v\n", n)
    r := f(n)
    watch("=> minimize returned value=%v\n", r)
    if r == 0 {
      break
    } else {
      n += 1
    }
  }
  return n
}




// ----- rf -----

func Divide(x, y int) int {
  watch("-> call Divide (%v,%v)\n", x, y)
  r := minimize(
    func(n int) int {
      return Subtract(increment(x), Multiply(y, increment(n)))
    })
  watch("=> Divide returned value=%v\n", r)
  return r
}




