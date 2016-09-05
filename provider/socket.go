package main

import (
  "net"
  "fmt"
  "bufio"
  //"strings"
  "os/exec"
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

  //sli := strings.Split(request, "/")
  //resp := exec.Command("php", "-s", "./services/php/" + sli[1] + "_impl.php " + sli[2] )

  //var resp interface{}
  //resp := exec.Command("php", "./provider/services/php/" + sli[1] + "_impl.php ", sli[2])
  resp := exec.Command("php", phpPath+"/user_impl.php", "ulist")
  //resp := exec.Command("php", "E:\\gitxx\\ttsoa\\test\\a.php")

  fmt.Println(resp)
  out, err := resp.CombinedOutput()
  checkError(err)
  fmt.Println(out)
  conn.Write([]byte(out))

}





