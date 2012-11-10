package gnet

import (
  "net"
  "encoding/binary"
  "time"
)

type Server struct {
  ln *net.TCPListener
  closed bool
  connPools map[uint64]*ConnPool
  New chan *Session
  newSessionBuffer *InfiniteSessionChan
  newConnChan *InfiniteTCPConnChan
  stop chan struct{}
  connPoolStopNotify *InfiniteConnPoolChan
}

func NewServer(addr string, key string) (*Server, error) {
  laddr, err := net.ResolveTCPAddr("tcp", addr)
  if err != nil {
    return nil, err
  }

  ln, err := net.ListenTCP("tcp", laddr)
  if err != nil {
    return nil, err
  }

  self := &Server{
    ln: ln,
    connPools: make(map[uint64]*ConnPool),
    New: make(chan *Session),
    newConnChan: NewInfiniteTCPConnChan(),
    stop: make(chan struct{}),
    connPoolStopNotify: NewInfiniteConnPoolChan(),
  }
  self.newSessionBuffer = NewInfiniteSessionChanWithOutChan(self.New)

  go self.startAcceptChan()
  go self.start(key)

  return self, nil
}

func (self *Server) startAcceptChan() {
  for {
    conn, err := self.ln.AcceptTCP()
    if err != nil {
      if self.closed {
        return
      }
      continue
    }
    self.newConnChan.In <- conn
  }
}

func (self *Server) start(key string) {
  heartBeat := time.Tick(time.Second * 2)
  tick := 0

  LOOP:
  for { // listen for incoming connection
    select {
    case conn := <-self.newConnChan.Out:
      var clientId uint64
      err := binary.Read(conn, binary.BigEndian, &clientId)
      if err != nil {
        continue
      }
      if self.connPools[clientId] == nil {
        self.log("new conn pool from %d", clientId)
        connPool := newConnPool(key, &(self.newSessionBuffer.In))
        connPool.clientId = clientId
        connPool.stopNotify = self.connPoolStopNotify.In
        self.connPools[clientId] = connPool
      }
      self.connPools[clientId].newConnChan.In <- conn

    case <-heartBeat:
      self.log("tick %d %d conn pools", tick, len(self.connPools))

    case connPool := <-self.connPoolStopNotify.Out:
      self.log("conn pool for client %d stop", connPool.clientId)
      delete(self.connPools, connPool.clientId)

    case <-self.stop:
      break LOOP
    }
    tick++
  }

  // finalizer
  for _, connPool := range self.connPools {
    connPool.Stop()
  }
  self.newSessionBuffer.Stop()
  self.newConnChan.Stop()
  self.connPoolStopNotify.Stop()
}

func (self *Server) Stop() {
  self.closed = true
  self.ln.Close()
  close(self.stop)
}

func (self *Server) log(f string, vars ...interface{}) {
  if DEBUG {
    colorp("35", "SERVER: " + f, vars...)
  }
}
