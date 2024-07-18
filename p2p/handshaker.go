package p2p

// HandshakerFunc
type HandshakerFunc func(any) error

func NOPHandshakefunc(any) error { return nil }
