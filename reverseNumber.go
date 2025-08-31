package main

import "fmt"
        
func main(){
    var n int 
    fmt.Print("Enter the value for print reverse order : ")
    fmt.Scan(&n)
    fmt.Print("the orginal number is",n,"its reveser format is ",reverse(n))
}
func reverse(n int) int{
    var rev int= 0
    for n > 0{
        digit := n%10
        rev = rev*10+digit
        n/=10
    }
    return rev
}

