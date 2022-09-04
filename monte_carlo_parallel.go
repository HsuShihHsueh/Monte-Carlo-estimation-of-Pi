package main

import (
	"fmt"
	"math"
	"math/rand"
)

const n = 1_000_000
var n_pi = 0.0

func average_value_method()(float64){
    n_pi = 0.0
    for i:=0; i<n; i++ {
        go average_step()
	}
	pi := 4 * n_pi / n
    return pi
}

func average_step(){
    tmp := rand.Float64()
    tmp = 1 - math.Pow(tmp, 2)
    tmp = math.Sqrt(tmp)
    n_pi += tmp 
}

func area_method()(float64){
    n_pi = 0.0
    for i:=0; i<n; i++ {
		go area_step()
    }
    pi := 4 * n_pi / n;
    return pi
}

func area_step(){
	tmp := math.Pow(rand.Float64(),2) + math.Pow(rand.Float64(),2)
	if tmp < 1 {
		n_pi++;
	}
}


func main() {
    result_average := average_value_method()
	fmt.Println("The estimate result of average value method: ",result_average)
	result_area := area_method()
	fmt.Println("The estimate result of average value method: ",result_area)
}