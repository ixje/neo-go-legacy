package consensus

import (
	coreb "github.com/ixje/neo-go-legacy/pkg/core/block"
	"github.com/ixje/neo-go-legacy/pkg/core/transaction"
	"github.com/ixje/neo-go-legacy/pkg/dbft/block"
	"github.com/ixje/neo-go-legacy/pkg/dbft/crypto"
	"github.com/ixje/neo-go-legacy/pkg/util"
)

// neoBlock is a wrapper of core.Block which implements
// methods necessary for dBFT library.
type neoBlock struct {
	coreb.Block

	signature []byte
}

var _ block.Block = (*neoBlock)(nil)

// Sign implements block.Block interface.
func (n *neoBlock) Sign(key crypto.PrivateKey) error {
	data := n.Base.GetHashableData()
	sig, err := key.Sign(data[:])
	if err != nil {
		return err
	}

	n.signature = sig

	return nil
}

// Verify implements block.Block interface.
func (n *neoBlock) Verify(key crypto.PublicKey, sign []byte) error {
	data := n.Base.GetHashableData()
	return key.Verify(data, sign)
}

// Transactions implements block.Block interface.
func (n *neoBlock) Transactions() []block.Transaction {
	txes := make([]block.Transaction, len(n.Block.Transactions))
	for i, tx := range n.Block.Transactions {
		txes[i] = tx
	}

	return txes
}

// SetTransactions implements block.Block interface.
func (n *neoBlock) SetTransactions(txes []block.Transaction) {
	n.Block.Transactions = make([]*transaction.Transaction, len(txes))
	for i, tx := range txes {
		n.Block.Transactions[i] = tx.(*transaction.Transaction)
	}
}

// Version implements block.Block interface.
func (n *neoBlock) Version() uint32 { return n.Block.Version }

// PrevHash implements block.Block interface.
func (n *neoBlock) PrevHash() util.Uint256 { return n.Block.PrevHash }

// MerkleRoot implements block.Block interface.
func (n *neoBlock) MerkleRoot() util.Uint256 { return n.Block.MerkleRoot }

// Timestamp implements block.Block interface.
func (n *neoBlock) Timestamp() uint64 { return uint64(n.Block.Timestamp) * 1000000000 }

// Index implements block.Block interface.
func (n *neoBlock) Index() uint32 { return n.Block.Index }

// ConsensusData implements block.Block interface.
func (n *neoBlock) ConsensusData() uint64 { return n.Block.ConsensusData }

// NextConsensus implements block.Block interface.
func (n *neoBlock) NextConsensus() util.Uint160 { return n.Block.NextConsensus }

// Signature implements block.Block interface.
func (n *neoBlock) Signature() []byte { return n.signature }
