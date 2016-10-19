package main

import "os/exec"

func execPool(adapter string, cmds string) string {
  cmd := exec.Command(adapter, cmds)
  out, err := cmd.CombinedOutput()
  checkError(err)
  return string(out)
}
