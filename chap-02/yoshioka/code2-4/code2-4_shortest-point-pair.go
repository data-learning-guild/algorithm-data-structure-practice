package main

import (
	"fmt"
	"math"
	"strconv"
)

//Function for calc distance
func calcDistance(x1, y1, x2, y2 float64) float64{
    return math.Sqrt(math.Pow((x1 - x2),2) + math.Pow((y1 - y2),2))
}

func main(){

    alldata := ReadCsv("points.csv")
    data := alldata[1:]

    n := len(data)

    x := make([]float64, n)
    y := make([]float64, n)

    for i := 0; i < n; i++ {
        f1, err := strconv.ParseFloat(data[i][0], 64)
        if err != nil {
            panic(err)
        }

        f2, err := strconv.ParseFloat(data[i][1], 64)
        if err != nil {
            panic(err)
        }

        x[i] = f1
        y[i] = f2
    }
    
    minimumDist := 1000000000.0

    // Start Search
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            // distance

            dist_i_j := calcDistance(x[i], y[i], x[j], y[j])
            
            if dist_i_j < minimumDist {
                minimumDist = dist_i_j
            }
        }
    }
    fmt.Printf("Result: %f\n", minimumDist)
}