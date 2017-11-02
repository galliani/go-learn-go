package main

import "fmt"

type Fund struct {
    // balance is unexported (private), because it's lowercase
    // So what we are doing here is similiar to declaring attributes that belong to Fund type
    balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
    // We can return a pointer to a new struct without worrying about
    // whether it's on the stack or heap: Go figures that out for us.
    return &Fund{
        balance: initialBalance,
    }
}

// Methods start with a *receiver*, in this case a Fund pointer
func (f *Fund) Balance() int {
    // Remember that f is an instance of Fund type Struct that does have balance attribute
    return f.balance
}

func (f *Fund) Withdraw(amount int) {
    f.balance -= amount
}


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

}