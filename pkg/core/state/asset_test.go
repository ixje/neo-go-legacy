package state

import (
	"testing"

	"github.com/ixje/neo-go-legacy/pkg/core/transaction"
	"github.com/ixje/neo-go-legacy/pkg/crypto/keys"
	"github.com/ixje/neo-go-legacy/pkg/internal/random"
	"github.com/ixje/neo-go-legacy/pkg/internal/testserdes"
	"github.com/ixje/neo-go-legacy/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeAssetState(t *testing.T) {
	asset := &Asset{
		ID:         random.Uint256(),
		AssetType:  transaction.Token,
		Name:       "super cool token",
		Amount:     util.Fixed8(1000000),
		Available:  util.Fixed8(100),
		Precision:  0,
		FeeMode:    feeMode,
		Owner:      keys.PublicKey{},
		Admin:      random.Uint160(),
		Issuer:     random.Uint160(),
		Expiration: 10,
		IsFrozen:   false,
	}

	testserdes.EncodeDecodeBinary(t, asset, new(Asset))
}

func TestAssetState_GetName_NEO(t *testing.T) {
	asset := &Asset{AssetType: transaction.GoverningToken}
	assert.Equal(t, "NEO", asset.GetName())
}

func TestAssetState_GetName_NEOGas(t *testing.T) {
	asset := &Asset{AssetType: transaction.UtilityToken}
	assert.Equal(t, "NEOGas", asset.GetName())
}
