package main

import "fmt"

func main(){
	var n int
    fmt.Scanln(&n)
    for i:=1; i <= n; i++{
        if (i % 3 == 0) && (i % 5 == 0){
            fmt.Println("FizzBuzz")
        }else if(i % 3 == 0) && (i % 5 != 0){
            fmt.Println("Fizz")
        }else if(i % 3 != 0) && (i % 5 != 0){
            fmt.Println(i)
        }
    }
}
