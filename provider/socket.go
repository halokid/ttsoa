package main

import (
  "net"
  "fmt"
  "bufio"
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
  request, err := reader.ReadString('\n')
  checkError(err)
  fmt.Println(request)

  //request like:  php:/user/ulist
  //resp := execPool("php", "E:\\gitxx\\ttsoa\\test\\a.php")

  sli := strings.Split(request, "/")
  resp := execPool(sli[0], "E:\\gitxx\\ttsoa\\test\\a.php")

  conn.Write([]byte(resp))

}





