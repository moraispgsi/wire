package fbp

import "errors"

const (
	connectionCapacity int = 10
)

type Sender interface {
	Send(IP interface{}) error
}

type Receiver interface {
	Recv() (interface{}, error)
}

type Connection interface {
	Send(IP interface{}) error
	Recv() (interface{}, error)
	SenderClosed() bool
	ReceiverClosed() bool
	CloseOutPort()
}

type connection struct {
	queue     chan interface{}
	inClosed  bool
	outClosed bool
}

func NewConnection() Connection {
	return &connection{
		queue:    make(chan interface{}, connectionCapacity),
		inClosed: false,
	}
}

func (conn *connection) Send(IP interface{}) error {
	if conn.inClosed {
		return errors.New("In has closed.")
	}
	conn.queue <- IP
	return nil
}

func (conn *connection) Recv() (interface{}, error) {
	IP, ok := <-conn.queue
	if !ok {
		conn.outClosed = true
		return nil, errors.New("Connection is closed.")
	}
	return IP, nil
}

func (conn *connection) SenderClosed() bool {
	return conn.inClosed
}

func (conn *connection) ReceiverClosed() bool {
	return conn.outClosed
}

func (conn *connection) CloseOutPort() {
	close(conn.queue)
}
