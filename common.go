package gnet

import (
  "fmt"
  "time"
)

const (
  CHAN_BUF_SIZE = 2 ^ 19

  PACKET_TYPE_SESSION = byte(0)
  PACKET_TYPE_PING = byte(1)

  SESSION_PACKET_TYPE_DATA = byte(0)
  SESSION_PACKET_TYPE_STATE = byte(1)
  SESSION_PACKET_TYPE_INFO = byte(2)

  STATE_FINISH_SEND = byte(0)
  STATE_ABORT_SEND = byte(1)
  STATE_FINISH_READ = byte(2)
  STATE_ABORT_READ = byte(3)

  NORMAL = iota
  FINISH
  ABORT

  IDLE_TIME_BEFORE_SESSION_CLOSE = time.Minute * 30
)

func p(f string, vars ...interface{}) {
  fmt.Printf(f, vars...)
}

func ps(f string, vars ...interface{}) string {
  return fmt.Sprintf(f, vars...)
}
