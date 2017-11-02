package main

import "fmt"

func main() {

    // Shorthand syntax of declaring and initializing vars
    // a := 10
    // b := "golang"
    // c := 4.17
    // d := true

    // Longform syntax
    // Declaration
    var a int
    var b string
    var c float64
    var d bool
    // Initializing
    a = 10
    b = "golang"
    c = 4.17
    d = true

    fmt.Printf("%v \n", a)
    fmt.Printf("%v \n", b)
    fmt.Printf("%v \n", c)
    fmt.Printf("%v \n", d)    
    
    fmt.Printf("%T \n", a)
    fmt.Printf("%T \n", b)
    fmt.Printf("%T \n", c)
    fmt.Printf("%T \n", d)
}