package main

import (
  "fmt"
  "os/exec")

func main() {
  cmd := exec.Command("php",  "F:\\GitHub\\ttsoa\\test\\a.php")
  out, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(out))
}

