package p2p

// HandshakerFunc
type HandshakerFunc func(Peer) error

func NOPHandshakefunc(Peer) error { return nil }
