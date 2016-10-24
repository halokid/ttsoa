package main

import (
  "net"
  "fmt"
  "bufio"
  //"strings"
  "os/exec"
  "strings"
)


func provider() {
  service := ":7777"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  checkError(err)
  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)
  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    go handleClient(conn)
  }
}



func handleClient(conn net.Conn) {
  fmt.Printf("from: %s\n", conn.RemoteAddr())
  defer conn.Close()

  //resp := execPool("php", "F:\\GitHub\\ttsoa\\test\\a.php")
  reader := bufio.NewReader(conn)
  //request, err := reader.ReadString('\n')
  request, err := reader.ReadString('\n')
  checkError(err)
  fmt.Println(request)

  //request like:  php:/user/ulist
  sli := strings.Split(request, "/")
  fmt.Println(sli)
  fmt.Println(sli[2])

  act := strings.Replace(sli[2], "\n", "", -1)
  fmt.Println(act)

  var resp *exec.Cmd
  if (sli[0] == "php") {
    //resp := exec.Command("php", phpPath+"/user_impl.php", "ulist")
    //resp := exec.Command("php", phpPath+"/user_impl.php", string(sli[2]))
    resp = exec.Command("php", phpPath + "/user_impl.php", act)
  } else if (sli[0] == "java") {
    resp = exec.Command("javac", javaPath + "xx.class", act)
  }

  fmt.Println(resp)
  out, err := resp.CombinedOutput()
  checkError(err)
  fmt.Println(out)
  conn.Write([]byte(out))

}





