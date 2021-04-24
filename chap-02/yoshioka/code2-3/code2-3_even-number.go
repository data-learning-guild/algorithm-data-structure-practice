package main

import "fmt"

func main(){
    var n int
    fmt.Print(">>> ")
    fmt.Scanf("%d", &n)

    for i := 2; i <= n; i+=2 {
        fmt.Printf("even number:%d\n", i)
    }

}