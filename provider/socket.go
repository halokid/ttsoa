package main

import (
  //"time"
  "net"
  //"bufio"
  //"fmt"
  //"time"
  "fmt"
  "bufio"
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
  fmt.Println("xxxxx")
  defer conn.Close()

  resp := execPool("php", "F:\\GitHub\\ttsoa\\test\\a.php")
  conn.Write([]byte(resp))

  /**
  //长连接读取方式
  reader := bufio.NewReader(conn)
  resp, err := reader.ReadString('\n')
  fmt.Println(resp)
  checkError(err)
  conn.Write([]byte("hello"))
  */
}
