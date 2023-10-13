package utils

import (
	"math"
)

func TrainTestSplit(x []float64, y []int64, split float64)([]float64, []float64, []int64 ,[]int64){
  l := len(x)
  splitNumber := split*float64(l)
  splitNumber = math.Floor(splitNumber)
  newSplit := int(splitNumber)

  x_train := []float64{}
  x_test := []float64{}
  y_train := []int64{}
  y_test := []int64{}
  
  for i:= 0; i<newSplit; i++{
    x_test = append(x_test, x[i])
    y_test = append(y_test, y[i])
  }
  x_train = x[newSplit:]
  y_train = y[newSplit:]
  return x_train, x_test, y_train, y_test
}
