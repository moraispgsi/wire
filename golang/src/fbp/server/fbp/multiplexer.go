package fbp

import "sync"

type MuxReceiver interface {
	Receiver
	AddReceiver(func() (interface{}, error))
	RemoveReceiver(ID int64)
	Listen()
}

//TODO: ADD mutex
type muxReceiver struct {
	connection         Connection
	receivers          map[int64]func() (interface{}, error)
	receiversCurrentID int64
	isListening        sync.Mutex
	canListen          sync.Mutex
}

func NewMuxReceiver() MuxReceiver {

	muxRecv := &muxReceiver{
		connection:         NewConnection(),
		receivers:          make(map[int64]func() (interface{}, error)),
		receiversCurrentID: 0,
		isListening:        sync.Mutex{},
		canListen:          sync.Mutex{},
	}

	return muxRecv
}

func (m *muxReceiver) listen(ID int64, wg *sync.WaitGroup) {
	defer wg.Done()

	recv, ok := m.receivers[ID]
	if !ok {
		return
	}

	for {

		IP, err := recv()
		if err != nil {
			return
		}

		err = m.connection.Send(IP)
		if err != nil {
			return
		}

	}

}

func (m *muxReceiver) Listen() {
	m.canListen.Lock()
	m.isListening.Lock()
	var wg sync.WaitGroup
	wg.Add(len(m.receivers))
	for ID := range m.receivers {
		go m.listen(ID, &wg)
	}
	go func() {
		wg.Wait()
		m.connection.CloseOutPort()
		m.isListening.Unlock()
	}()
}

func (m *muxReceiver) AddReceiver(receiver func() (interface{}, error)) {
	m.canListen.Lock()
	ID := m.receiversCurrentID
	m.receivers[ID] = receiver
	m.receiversCurrentID++
	m.canListen.Unlock()
}

func (m *muxReceiver) RemoveReceiver(ID int64) {
	delete(m.receivers, ID)
}

func (m *muxReceiver) Recv() (interface{}, error) {
	IP, err := m.connection.Recv()
	if err != nil {
		m.canListen.Unlock()
	}
	return IP, err
}

func (m *muxReceiver) ReceiverClosed() bool {
	m.canListen.Lock()
	defer m.canListen.Unlock()
	return m.connection.ReceiverClosed()
}
