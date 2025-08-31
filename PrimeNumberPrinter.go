package main

import ( "fmt"
        "math"
        )

func main(){
    var n,i int 
    fmt.Print("Enter the value for print prime number ")
    fmt.Scan(&n)
    for i = 2 ; i <= n ;i++{
        if isPrime(i){
            fmt.Print(" ",i)
        }
    }
}
func isPrime(n int)bool{
    if n < 2{
        return false
    }
    if n== 2{
        return true
    }
    if n % 2 == 0{
        return false
    }
    for i := 3 ; i <= int(math.Sqrt(float64(n)));i+=2 {
        if n%i == 0{
            return false
        }
    } 
    return true
}
