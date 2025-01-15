package payload

import (
	"testing"

	"github.com/ixje/neo-go-legacy/pkg/crypto/hash"
	"github.com/ixje/neo-go-legacy/pkg/internal/testserdes"
	. "github.com/ixje/neo-go-legacy/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestInventoryEncodeDecode(t *testing.T) {
	hashes := []Uint256{
		hash.Sha256([]byte("a")),
		hash.Sha256([]byte("b")),
	}
	inv := NewInventory(BlockType, hashes)

	testserdes.EncodeDecodeBinary(t, inv, new(Inventory))
}

func TestEmptyInv(t *testing.T) {
	msgInv := NewInventory(TXType, []Uint256{})

	data, err := testserdes.EncodeBinary(msgInv)
	assert.Nil(t, err)
	assert.Equal(t, []byte{byte(TXType), 0}, data)
	assert.Equal(t, 0, len(msgInv.Hashes))
}
