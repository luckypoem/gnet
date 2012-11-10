package gnet

import (
  "fmt"
)

const (
  INITIAL_BUF_CAPACITY = 2 ^ 16

  PACKET_TYPE_SESSION = byte(0)
  PACKET_TYPE_PING = byte(1)
  PACKET_TYPE_INFO = byte(2)

  SESSION_PACKET_TYPE_DATA = byte(0)
  SESSION_PACKET_TYPE_STATE = byte(1)
  SESSION_PACKET_TYPE_INFO = byte(2)

  STATE_FINISH_SEND = byte(0)
  STATE_ABORT_SEND = byte(1)
  STATE_FINISH_READ = byte(2)
  STATE_ABORT_READ = byte(3)
  STATE_STOP = byte(4)

  NORMAL = iota
  FINISH
  ABORT

  DATA
  STATE
  INFO
)

func p(f string, vars ...interface{}) {
  fmt.Printf(f, vars...)
}

func ps(f string, vars ...interface{}) string {
  return fmt.Sprintf(f, vars...)
}

func colorp(color string, s string, vars... interface{}) {
  println("\033[" + color + "m" + fmt.Sprintf(s, vars...) + "\033[0m")
}
