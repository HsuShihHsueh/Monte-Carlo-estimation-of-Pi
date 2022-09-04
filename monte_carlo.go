package main

import (
	"fmt"
	"math"
	"math/rand"
)

const n = 1000000

func average_value_method()(float64){
    pi := 0.0
    for i:=0; i<n; i++ {
        tmp := rand.Float64()
        tmp = 1 - math.Pow(tmp, 2)
        tmp = math.Sqrt(tmp)
        pi += tmp
	}
	pi = 4 * pi / n;
    return pi
}

func area_method()(float64){
    pi := 0.0
    for i:=0; i<n; i++ {
        tmp := math.Pow(rand.Float64(),2) + math.Pow(rand.Float64(),2)
        if tmp < 1 {
            pi++;
        }
    }
    pi = 4 * pi / n;
    return pi
}

func main() {
    result_average := average_value_method()
	fmt.Println("The estimate result of average value method: ",result_average)
	result_area := area_method()
	fmt.Println("The estimate result of average value method: ",result_area)
}
