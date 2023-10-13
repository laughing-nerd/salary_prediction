package utils

func Predict(x []float64, m float64, c float64) []float64{

  y_predict := []float64{}
  for _,v := range x{
    yi := (m * v) + c
    y_predict = append(y_predict, yi)
  }

  return y_predict
}
