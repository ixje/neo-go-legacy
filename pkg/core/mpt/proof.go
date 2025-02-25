package mpt

import (
	"bytes"

	"github.com/ixje/neo-go-legacy/pkg/core/storage"
	"github.com/ixje/neo-go-legacy/pkg/crypto/hash"
	"github.com/ixje/neo-go-legacy/pkg/util"
)

// GetProof returns a proof that key belongs to t.
// Proof consist of serialized nodes occuring on path from the root to the leaf of key.
func (t *Trie) GetProof(key []byte) ([][]byte, error) {
	var proof [][]byte
	path := toNibbles(key)
	r, err := t.getProof(t.root, path, &proof)
	if err != nil {
		return proof, err
	}
	t.root = r
	return proof, nil
}

func (t *Trie) getProof(curr Node, path []byte, proofs *[][]byte) (Node, error) {
	switch n := curr.(type) {
	case *LeafNode:
		if len(path) == 0 {
			*proofs = append(*proofs, copySlice(n.Bytes()))
			return n, nil
		}
	case *BranchNode:
		*proofs = append(*proofs, copySlice(n.Bytes()))
		i, path := splitPath(path)
		r, err := t.getProof(n.Children[i], path, proofs)
		if err != nil {
			return nil, err
		}
		n.Children[i] = r
		return n, nil
	case *ExtensionNode:
		if bytes.HasPrefix(path, n.key) {
			*proofs = append(*proofs, copySlice(n.Bytes()))
			r, err := t.getProof(n.next, path[len(n.key):], proofs)
			if err != nil {
				return nil, err
			}
			n.next = r
			return n, nil
		}
	case *HashNode:
		if !n.IsEmpty() {
			r, err := t.getFromStore(n.Hash())
			if err != nil {
				return nil, err
			}
			return t.getProof(r, path, proofs)
		}
	}
	return nil, ErrNotFound
}

// VerifyProof verifies that path indeed belongs to a MPT with the specified root hash.
// It also returns value for the key.
func VerifyProof(rh util.Uint256, key []byte, proofs [][]byte) ([]byte, bool) {
	path := toNibbles(key)
	tr := NewTrie(NewHashNode(rh), false, storage.NewMemCachedStore(storage.NewMemoryStore()))
	for i := range proofs {
		h := hash.DoubleSha256(proofs[i])
		// no errors in Put to memory store
		_ = tr.Store.Put(makeStorageKey(h[:]), proofs[i])
	}
	_, bs, err := tr.getWithPath(tr.root, path)
	return bs, err == nil
}
