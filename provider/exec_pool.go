package main

import "os/exec"

func execPool(cmm string) string {
  cmd := exec.Command(cmm)
  out, err := cmd.CombinedOutput()
  checkError(err)
  return string(out)
}
