package main

import (
  "net"
  "fmt"
<<<<<<< HEAD
=======
  //"bufio"
  //"bufio"
>>>>>>> 73f2bbf4225e42332e1d6f063a8322fc51aa1b71
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
  fmt.Printf("from: %s\n", conn.RemoteAddr())
  defer conn.Close()

  //resp := execPool("php", "F:\\GitHub\\ttsoa\\test\\a.php")
  reader := bufio.NewReader(conn)
  request, err := reader.ReadString('\n')
  checkError(err)
  fmt.Println(request)

  resp := execPool("php", "E:\\gitxx\\ttsoa\\test\\a.php")
  conn.Write([]byte(resp))

}





