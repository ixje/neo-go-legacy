package result

import (
	"github.com/ixje/neo-go-legacy/pkg/crypto/keys"
	"github.com/ixje/neo-go-legacy/pkg/util"
)

// Validator used for the representation of
// state.Validator on the RPC Server.
type Validator struct {
	PublicKey keys.PublicKey `json:"publickey"`
	Votes     util.Fixed8    `json:"votes"`
	Active    bool           `json:"active"`
}
