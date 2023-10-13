package utils

var LearningRate float64 = 0.0001

func LinearRegression(x_train []float64, y_train []int64)(float64, float64){
  l := len(x_train)

  // Least Squares Method
  m:=0.0
  c:=0.0
  var (
    sumXY float64
    sumX float64
    sumY float64
    sumX2 float64
  )
  for i:=0;i<l;i++{
    sumX = sumX + x_train[i]
    sumY = sumY + float64(y_train[i])
    sumXY = sumXY + (x_train[i] * float64(y_train[i]))
    sumX2 = sumX2 + (x_train[i] * x_train[i])
  }
  m = ((float64(l) * sumXY) - (sumX * sumY)) / ((float64(l)*sumX2) - (sumX*sumX))
  c = (sumY - (m*sumX)) / (float64(l))

  //Gradient Descent algorithm to find the closest m and c
  // m:=0.0
  // c:=0.0
  // epoch := 10000000000
  // for i:=0;i<epoch;i++{
  //   var (
  //     sumM float64
  //     sumC float64
  //   )
  //
  //   for j:=0;j<l;j++{
  //     sumM = sumM + (x_train[j] * (float64(y_train[j]) - ((m*x_train[j])+c)))
  //     sumC = sumC + float64(y_train[j]) - ((m*x_train[j])+c)
  //   }
  //   dM := -(1.0/float64(l)) * sumM
  //   dC := -(1.0/float64(l)) * sumC
  //
  //   m = m - (LearningRate * dM)
  //   c = c - (LearningRate * dC)
  // }
  return m,c
}
