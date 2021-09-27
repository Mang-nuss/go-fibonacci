package main

import (
  "fmt"
  //"log"
  //"io/ioutil"
  //"strconv"
)

func fibonacci(nr int) (res []int, err error) {
  result := []int{1}

  if err != nil {
    fmt.Println("there was an error in the fibonacci function")
  } else {
    for i := 0; i <= nr; i++ {
      result = []int{i} // use := instead of var in the beginning
    }
  }

  return result, err
}

func fibonacciRecursive(nr int, i int, sequence []int) (res []int, err error) {

  f := i
  if err != nil {
    fmt.Println("there was an error")
  } else {
    //
    // n := 0
    index := i
    for index <= nr {
      if f > nr-1 {
        fmt.Println("f =", f, "-> break")
        break

      if f == 0 {
        sequence[f] = 0
      } else if f == 1 {
        sequence[f] = 1
      } else if f == 2 {
        sequence[f] = 1
      } else {
        sequence[f] = sequence[f-1]+sequence[f-2]
      }
      fmt.Println("sequence:", sequence)

      fibonacciRecursive(index, i+1, sequence)
      f += 1
      fmt.Println("f =", f)
      //fmt.Println("i =", i)
    }
    //index += 1
  }

  fmt.Println("outside for loop,", )
  //result := []int{nr}
  return sequence, err
}

func main() {

  nr := 3
  seq := make([]int, nr) //this is how to initialize an array of length nr
  fib, err := fibonacciRecursive(nr, 0, seq)

  if err != nil {
    fmt.Println("there was an error in the main function")
  } else {

    i := 0
    fmt.Println("Main outprint:")

    for i < nr { //range: extension
      fmt.Println(fib[i])
      i += 1
    }
  }
}
