package main

import "fmt"

func main(){
  fmt.Println("Hello World")
  
  var age int = 40
  var favNum float64 = 89.83
  
  randNum := 1
  
  fmt.Println(age,favNum,randNum)
  // randNum = "Hello" cause an error

  // /////////////////part 2
  const pi float64 = 3.14159265
  
  // var (
  //   varA = 2
  //   varB = 3  // const pi float64 = 3.14159265
  
  var (
    varA = 2
    varB = 3
  )
  
  var myName string = "Javier Lona"
  fmt.Println(len(myName))
  fmt.Println(varA,varB)
  
  fmt.Println(myName + " is a robot")
  
  var isOver40 bool = true
  fmt.Println("Print the bool ", isOver40)
  fmt.Printf("%.3f \n", pi)
  fmt.Printf("%T %S \n", pi," The variable pi is this ")
  fmt.Printf("%b \n", 8)

  i := 1
  for i <= 10 {
    fmt.Println(i)
    i++
  }
}
