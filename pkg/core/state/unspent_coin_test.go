package state

import (
	"testing"

	"github.com/ixje/neo-go-legacy/pkg/core/transaction"
	"github.com/ixje/neo-go-legacy/pkg/internal/random"
	"github.com/ixje/neo-go-legacy/pkg/internal/testserdes"
	"github.com/ixje/neo-go-legacy/pkg/util"
)

func TestDecodeEncodeUnspentCoin(t *testing.T) {
	unspent := &UnspentCoin{
		Height: 100500,
		States: []OutputState{
			{
				Output: transaction.Output{
					AssetID:    random.Uint256(),
					Amount:     util.Fixed8(42),
					ScriptHash: random.Uint160(),
				},
				SpendHeight: 201000,
				State:       CoinSpent,
			},
			{
				Output: transaction.Output{
					AssetID:    random.Uint256(),
					Amount:     util.Fixed8(420),
					ScriptHash: random.Uint160(),
					Position:   1,
				},
				SpendHeight: 0,
				State:       CoinConfirmed,
			},
			{
				Output: transaction.Output{
					AssetID:    random.Uint256(),
					Amount:     util.Fixed8(4200),
					ScriptHash: random.Uint160(),
					Position:   2,
				},
				SpendHeight: 111000,
				State:       CoinSpent & CoinClaimed,
			},
		},
	}

	testserdes.EncodeDecodeBinary(t, unspent, new(UnspentCoin))
}
