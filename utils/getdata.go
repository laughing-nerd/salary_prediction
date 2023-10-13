package utils

import (
  "os"
  "encoding/csv"
)

func GetData(filename string)[][]string{
  fd, err := os.Open(filename)
  if err!=nil{
    panic(err)
  }
  defer fd.Close()
  reader :=csv.NewReader(fd)
  contents, read_err := reader.ReadAll()
  if read_err!=nil{
    panic(read_err)
  }
  return contents
}
