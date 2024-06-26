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

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
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
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf("new incoming connection from %s\n", conn)
}
