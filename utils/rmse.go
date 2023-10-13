package utils

import (
	"fmt"
	"math"
)

func RMSE(predicted []float64, actual []int64){
  l := len(predicted)

  var sum float64

  for i:=0; i<l; i++ {
    difference := float64(actual[i]) - predicted[i]
    squaredDifference := difference * difference
    sum += squaredDifference
  }
  mse := sum / float64(l)
  fmt.Println("RMSE = ", math.Sqrt(mse))
}
