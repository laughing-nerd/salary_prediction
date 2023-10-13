package main

import(
  "fmt"
  "strconv"
  utils "github.com/laughing-nerd/salary-prediction/utils"
)

func main(){
  contents := utils.GetData("./data/Salary_Data.csv") //Get the contents of the CSV file

  var x []float64;
  var y []int64;

  // Separate x and y from CSV
  for i:=1;i<len(contents);i++{
    x_float, _ := strconv.ParseFloat(contents[i][0], 64)
    y_int, _ := strconv.ParseInt(contents[i][1], 10, 64)
    x = append(x, x_float)
    y = append(y, y_int)
  }
  x_train, x_test, y_train, y_test := utils.TrainTestSplit(x, y, 0.2) //Train test split

  m,c := utils.LinearRegression(x_train, y_train)
  fmt.Println("m = ", m)
  fmt.Println("c = ", c)
  fmt.Println("Salary for 5 YOE = ", (m*5)+c) //Salary for 5 YOE
  y_predict := utils.Predict(x_test, m, c)

  fmt.Println("Predicted = ", y_predict)
  fmt.Println("Actual = ", y_test)

  utils.RMSE(y_predict, y_test)
}
