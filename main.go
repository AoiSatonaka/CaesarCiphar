package main

import (
  "flag"
  "fmt"
  "os"
)

func main() {
  const origin = "0123456789"
  const shift = "1234567890"

  flag.Parse()
  //f,err := os.Open(flag.Arg(1))
  //if err!=nil {
  //  fmt.Print("error")
  //}
  //defer f.Close()

  //第一引数に暗号化したい文字列を渡す
  for _,r := range flag.Arg(0) {
    
  }

}
