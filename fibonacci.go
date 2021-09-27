package main

import (
  "fmt"
  //"log"
  //"io/ioutil"
  //"strconv"
)

func fibonacci(nr int) (res []int, err error) {
  //var n = nr
  //fib := [n]int
  //var fib []int = make([]int, nr)
  result := []int{1}

  if err != nil {
    fmt.Println("there was an error in the fibonacci function")
  } else { //does this need to be printed on the same line?
    // result := []int{1}
    for i := 0; i <= nr; i++ {
      //fmt.Printf(i, "\n")
      result = []int{i} // use := instead of var in the beginning
    }
  }

  // result := []int{1}
  // for i := 0; i < nr; i++ {
  //   //fmt.Printf(i, "\n")
  //   result = []int{i} // use := instead of var in the beginning
  // }

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
    //fmt.Println("index before for loop=", index)
    //fmt.Println("nr =", f)
    for index <= nr {
      //if f < 3 { break }
      //fmt.Println("index =", index)
      // for n <= nr {
      //   fmt.Println(n, ",")
      // }
      //f += index
      if f > nr-1 {
        fmt.Println("f =", f, "-> break")
        break
      }
      // if index > 3 {
      //   fmt.Println("index: break")
      //   break
      // }
      //var length = nr
      //sequence := make([]int, nr)
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

  //FOR THE INITIAL FUNCTION
  // fmt.Println("the numbers: ")
  // result, err := fibonacci(4)
  //
  // if err != nil {
  //   fmt.Println("there was an error in the main function")
  // } else {
  //   for i := range result { //range: extension
  //   fmt.Println(result[i])
  //   }
  // }
}
