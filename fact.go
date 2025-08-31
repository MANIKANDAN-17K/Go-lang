package main

import "fmt"

func main(){
    fmt.Print("Enter the value for factorial : ")
    var n int
    fmt.Scan(&n)
    var fact int = 1
    for i := 1;i<=n;i++{
        fact *= i
    }
    fmt.Println("The factorial value of ",n," is",fact)
}
