package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  //const origin = "0123456789"
  //const shift = "1234567890"

  fmt.Print("暗号化したい文字列を入力してください。：")
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  plainRuneArray := []rune(stdin.Text())

  var shiftRuneArray []rune
  for _,rune := range plainRuneArray {

  }
  fmt.Print(plainRuneArray)

}