package gnet

import (
  "fmt"
)

const (
  CHAN_BUF_SIZE = 2 ^ 19
)

func p(f string, vars ...interface{}) {
  fmt.Printf(f, vars...)
}
