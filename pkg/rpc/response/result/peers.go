package result

import (
	"strconv"
	"strings"
)

type (
	// GetPeers payload for outputting peers in `getpeers` RPC call.
	GetPeers struct {
		Unconnected Peers `json:"unconnected"`
		Connected   Peers `json:"connected"`
		Bad         Peers `json:"bad"`
	}

	// Peers represent a slice of peers.
	Peers []Peer

	// Peer represents the peer.
	Peer struct {
		Address string `json:"address"`
		Port    uint16 `json:"port"`
	}
)

// NewGetPeers creates a new GetPeers structure.
func NewGetPeers() GetPeers {
	return GetPeers{
		Unconnected: []Peer{},
		Connected:   []Peer{},
		Bad:         []Peer{},
	}
}

// AddUnconnected adds a set of peers to the unconnected peers slice.
func (g *GetPeers) AddUnconnected(addrs []string) {
	g.Unconnected.addPeers(addrs)
}

// AddConnected adds a set of peers to the connected peers slice.
func (g *GetPeers) AddConnected(addrs []string) {
	g.Connected.addPeers(addrs)
}

// AddBad adds a set of peers to the bad peers slice.
func (g *GetPeers) AddBad(addrs []string) {
	g.Bad.addPeers(addrs)
}

// addPeers adds a set of peers to the given peer slice.
func (p *Peers) addPeers(addrs []string) {
	for i := range addrs {
		addressParts := strings.Split(addrs[i], ":")
		port, _ := strconv.Atoi(addressParts[1]) // We know it's a good port number.
		peer := Peer{
			Address: addressParts[0],
			Port:    uint16(port),
		}

		*p = append(*p, peer)
	}
}
