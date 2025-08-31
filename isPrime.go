package main

import ("fmt"

        "math"
)


func main(){
    var n int 
    fmt.Print("Enter the value for check the number is Prime or not : ")
    fmt.Scan(&n) 
    
    if isPrime(n){
        fmt.Print("the given number",n," is prime ")
    }else{
        fmt.Print("the given number",n," is not prime")
    }
}

func isPrime(n int)bool {
    if n < 2{
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    for i:=3; i<= int(math.Sqrt(float64(n))); i+=2 {
        if n % i == 0{
            return false
        }
    }
    
    return true
}
