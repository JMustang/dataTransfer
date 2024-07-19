package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPeer representa o nó remoto através de uma conexão TCP estabelecida
type TCPPeer struct {
	// conn é a conexão subjacente do peer
	conn net.Conn

	// se discarmos e recuperarmos uma conexão => outbound == true
	// se aceitar e recuperar uma conexão => saída == falso
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakerFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	go t.startAcceptloop()

	return nil
}

func (t *TCPTransport) startAcceptloop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		fmt.Printf("new incoming connection from %+v\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {

		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}

	// lenDecodeError := 0
	// Read loop
	msg := &Temp{}
	for {
		if err := t.Decoder.Decoder(conn, msg); err != nil {
			fmt.Printf("TCP decode error: %s\n", err)
			continue
			// lenDecodeError++
			// if lenDecodeError > 5 {
			// 	break
			// }
		}
	}

}
