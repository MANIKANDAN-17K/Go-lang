package main

import "fmt"

func main(){
    var n int 
    fmt.Print("Enter the integer for check palindrome or not  : ")
    fmt.Scan(&n)
    
    if isPalindrome(n){
        fmt.Print("the given number ",n," is palindrome")
    }else{
        fmt.Print("the given number ",n," is not palindrome")
    }
}

func isPalindrome(n int) bool{
    divisor := 1 
    for n/divisor > 10 {
        divisor = divisor * 10
    }
    
    for n != 0 {
        left := n/divisor
        right := n%10 
        if left != right{
            return false
        }
        n = (n%divisor) / 10
        divisor = divisor / 100
    }
    return true
}
