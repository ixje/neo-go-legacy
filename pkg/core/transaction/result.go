package transaction

import "github.com/ixje/neo-go-legacy/pkg/util"

// Result represents the Result of a transaction.
type Result struct {
	AssetID util.Uint256
	Amount  util.Fixed8
}
