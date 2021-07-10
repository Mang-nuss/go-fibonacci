package main

import (
  "fmt"
  "testing"
)

// Equals tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equals(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }

    fmt.Println("the slices are equal")
    return true
}

func FibonacciEquals(val1 int, val2 []int) bool {
  val, err := fibonacci(val1)

  if err != nil {
    fmt.Println("error in Fibonacci calculation")
    return false

  } else {
    //fmt.Println(val, " is equal to ", val2)
    return Equals(val, val2)
  }
}

func TestFibonacci(t *testing.T) { //the method name start must be capital letter

  //value, err := fibonacci(1)

  // if err != nil {
  //   t.Fatalf("error in Fibonacci calculation")
  // }

  value := 10
  seq := make([]int, value)
  expected := []int{0,1,1,2,3,5,8,13,21,34}
  fib, err := fibonacciRecursive(value, 0, seq)
  fmt.Println("fibonacci tells", fib)

  if err != nil {
    fmt.Println("error in Fibonacci calculation")
  }

  if !Equals(fib, expected) {
    //t.Fatalf("the value - %c", value) //%c = formatting directive for integer
    t.Fatal("the sequence", fib, "is not equal to", expected) //no spaces are needed
  }
}
