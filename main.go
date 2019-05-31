package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  for {
    //input
    plainRuneArray := []rune(getInput("暗号化したい文字列を入力してください"))
    shifter, err := strconv.Atoi(getInput("増分を指定してください"))
    if err != nil {
      fmt.Println("入力値が不正です。最初からやりなおしてください")
      continue
    }

    fmt.Println("平文(rune)：",plainRuneArray)

    //encryption
    var encRuneArray []rune
    for _, rune := range plainRuneArray {
      //golang rune is an alias for int32
      encRuneArray = append(encRuneArray,rune+int32(shifter))
    }

    fmt.Println("暗号文(rune)：",encRuneArray)
    fmt.Println("暗号文：",string(encRuneArray))

    //decryption
    var decRuneArray []rune
    for _,rune := range encRuneArray {
      decRuneArray = append(decRuneArray,rune-int32(shifter))
    }

    fmt.Println("復号した平文(rune)：",decRuneArray)
    fmt.Println("復号した平文：",string(decRuneArray))


    if getInput("続行=>[Enter] 終了=>[e+Enter]") == "e"{
      break
    }
  }
}

func getInput(tag string) string {
  fmt.Print(tag+"：")
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  return stdin.Text()
}