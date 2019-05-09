package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")

    var i int 
    var j int
    j= 2
    var myfloat float64
    var o int = 1

    fmt.Println(i) // prints out -127
    fmt.Println(j) // prints out -127
    fmt.Println(o) // prints out -127


    var mynewint int
    myfloat = float64(mynewint)

    fmt.Println(myfloat) 

    var myfloat3 float64
myint2 := int(myfloat3)
fmt.Println(myint2) // prints out -127





    var myint int8
    for i := 0; i < 129; i++ {
        myint += 1
        fmt.Println(myint) // prints out -127

    }
  //  fmt.Println(myint) // prints out -127

  var isTrue bool = true
  var isFalse bool = false
  // AND
  if isTrue && isFalse {
    fmt.Println("Both Conditions need to be True")
  }
  // OR - not exclusive
  if isTrue || isFalse {
    fmt.Println("Only one condition needs to be True")
  }





}