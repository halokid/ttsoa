package main

import (
  "fmt"
  "os/exec")

func main() {
  //cmd := exec.Command("php",  "F:\\GitHub\\ttsoa\\test\\a.php")

  //php
  //cmd := exec.Command("php", "E:\\gitxx\\ttsoa\\test\\a.php")

  //java
  //_ = exec.Command("cd", "E:\\gitxx\\ttsoa\\test")
  //cmd2 := exec.Command("java", "JavaTest")
  cmd2 := exec.Command("java")

  //out, err := cmd.CombinedOutput()
  out, err := cmd2.CombinedOutput()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(out))
}

