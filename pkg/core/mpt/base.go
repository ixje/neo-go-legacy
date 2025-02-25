package mpt

import (
	"fmt"

	"github.com/ixje/neo-go-legacy/pkg/crypto/hash"
	"github.com/ixje/neo-go-legacy/pkg/io"
	"github.com/ixje/neo-go-legacy/pkg/util"
)

// BaseNode implements basic things every node needs like caching hash and
// serialized representation. It's a basic node building block intended to be
// included into all node types.
type BaseNode struct {
	hash       util.Uint256
	bytes      []byte
	hashValid  bool
	bytesValid bool
}

// BaseNodeIface abstracts away basic Node functions.
type BaseNodeIface interface {
	Hash() util.Uint256
	Type() NodeType
	Bytes() []byte
}

type flushedNode interface {
	setCache([]byte, util.Uint256)
}

func (b *BaseNode) setCache(bs []byte, h util.Uint256) {
	b.bytes = bs
	b.hash = h
	b.bytesValid = true
	b.hashValid = true
}

// getHash returns a hash of this BaseNode.
func (b *BaseNode) getHash(n Node) util.Uint256 {
	if !b.hashValid {
		b.updateHash(n)
	}
	return b.hash
}

// getBytes returns a slice of bytes representing this node.
func (b *BaseNode) getBytes(n Node) []byte {
	if !b.bytesValid {
		b.updateBytes(n)
	}
	return b.bytes
}

// updateHash updates hash field for this BaseNode.
func (b *BaseNode) updateHash(n Node) {
	if n.Type() == HashT {
		panic("can't update hash for hash node")
	}
	b.hash = hash.DoubleSha256(b.getBytes(n))
	b.hashValid = true
}

// updateCache updates hash and bytes fields for this BaseNode.
func (b *BaseNode) updateBytes(n Node) {
	buf := io.NewBufBinWriter()
	encodeNodeWithType(n, buf.BinWriter)
	b.bytes = buf.Bytes()
	b.bytesValid = true
}

// invalidateCache sets all cache fields to invalid state.
func (b *BaseNode) invalidateCache() {
	b.bytesValid = false
	b.hashValid = false
}

// encodeNodeWithType encodes node together with it's type.
func encodeNodeWithType(n Node, w *io.BinWriter) {
	w.WriteB(byte(n.Type()))
	n.EncodeBinary(w)
}

// DecodeNodeWithType decodes node together with it's type.
func DecodeNodeWithType(r *io.BinReader) Node {
	if r.Err != nil {
		return nil
	}
	var n Node
	switch typ := NodeType(r.ReadB()); typ {
	case BranchT:
		n = new(BranchNode)
	case ExtensionT:
		n = new(ExtensionNode)
	case HashT:
		n = new(HashNode)
	case LeafT:
		n = new(LeafNode)
	default:
		r.Err = fmt.Errorf("invalid node type: %x", typ)
		return nil
	}
	n.DecodeBinary(r)
	return n
}
