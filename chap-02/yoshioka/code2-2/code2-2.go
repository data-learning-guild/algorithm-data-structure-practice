package main

import "fmt"

func main(){
    var n int
    fmt.Print(">>> ")
    fmt.Scanf("%d", &n)

    count := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            count++
            fmt.Printf("count: %d\n", count)                
        }
    } 
}