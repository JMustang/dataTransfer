package p2p

// Peer é uma interface que representa o nó remoto.
type Peer interface {
}

// Transporte é tudo que cuida da comunicação
// entre os nós da rede. Isto pode ser do
// formulário (TCP, UDP, websockets, ...)
type Transporte interface {
	ListenAndAccept() error
}

// Network é a rede de nós.
