package main

import "fmt"

func main(){
    var n,i int
    fmt.Print("Enter the value for find the  fibonacci series : ")
    fmt.Scan(&n)
    
    a := 0
    b := 1
    fmt.Print(" ",a)
    for i=0;i<n;i++{
        c := a + b
        a = b 
        b = c 
        fmt.Print(" ",a)
    }
}
