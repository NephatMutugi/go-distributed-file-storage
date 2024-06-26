package p2p

// Peer is an interface that represents the remote nodes.
type Peer interface {
}

// Transport is Anything that handles communication
// between nodes in the network. This can be of the
// form (TCP, UDP, websockets...)
type Transport interface {
}
