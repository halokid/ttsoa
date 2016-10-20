package main

import (
  "fmt"
  "strings"
)

func main() {

  str := "php/user/ulist"
  fmt.Println(str)

  sli := strings.Split(str, "/");
  fmt.Println(sli)
}
