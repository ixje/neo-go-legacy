package payload

import (
	"testing"

	"github.com/ixje/neo-go-legacy/pkg/crypto/hash"
	"github.com/ixje/neo-go-legacy/pkg/internal/testserdes"
	"github.com/ixje/neo-go-legacy/pkg/util"
)

func TestGetBlockEncodeDecode(t *testing.T) {
	start := []util.Uint256{
		hash.Sha256([]byte("a")),
		hash.Sha256([]byte("b")),
		hash.Sha256([]byte("c")),
		hash.Sha256([]byte("d")),
	}

	p := NewGetBlocks(start, util.Uint256{})
	testserdes.EncodeDecodeBinary(t, p, new(GetBlocks))
}

func TestGetBlockEncodeDecodeWithHashStop(t *testing.T) {
	var (
		start = []util.Uint256{
			hash.Sha256([]byte("a")),
			hash.Sha256([]byte("b")),
			hash.Sha256([]byte("c")),
			hash.Sha256([]byte("d")),
		}
		stop = hash.Sha256([]byte("e"))
	)
	p := NewGetBlocks(start, stop)
	testserdes.EncodeDecodeBinary(t, p, new(GetBlocks))
}
