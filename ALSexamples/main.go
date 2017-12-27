package main

import (
	"fmt"
	. "github.com/mojocn/recommend-go/bayesianFilter"
)

func main() {

	mat := MakeRatingMatrix([]float64{
		4, NA, 5, 5,
		4, 2, 1, NA,
		3, NA, 2, 4,
		4, 4, NA, NA,
		2, 1, 3, 5}, 5, 4)

	//
	// do something with err
	fmt.Println(BayesianFilter(mat, 0, 1))

}
