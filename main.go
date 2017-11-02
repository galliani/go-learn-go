package main

import "fmt"

// Since we put all the Fund related functions in a separate file, 
// but still in same package "main", to validly run this script we have to type either:
// go run *.go
//      or
// go build
// then -> ./hello_go
// 
// After that, we should clean the exe bin by typing "go clean" in terminal
func main() {
    fund := NewFund(100)
    fmt.Printf("%v \n", fund)
    fmt.Printf("%T \n", fund)

    for i := 0; i < 101; i++ {
        fund.Withdraw(1)
    }

    if fund.Balance() > 0 {
        fmt.Println("You still have money, yeay!")
    } else {
        fmt.Println("Damn, you have got no money left!")
    }

}

// Just some notes about variable in go:
// Shorthand syntax of declaring and initializing vars
// a := 10
// b := "golang"
// c := 4.17
// d := true

// // Longform syntax
// // Declaration
// var a int
// var b string
// var c float64
// var d bool
// // Initializing
// a = 10
// b = "golang"
// c = 4.17
// d = true

// fmt.Printf("%v \n", a)
// fmt.Printf("%v \n", b)
// fmt.Printf("%v \n", c)
// fmt.Printf("%v \n", d)    

// fmt.Printf("%T \n", a)
// fmt.Printf("%T \n", b)
// fmt.Printf("%T \n", c)
// fmt.Printf("%T \n", d)