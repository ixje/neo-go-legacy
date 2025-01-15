package transaction

import "github.com/ixje/neo-go-legacy/pkg/io"

// TXer is interface that can act as the underlying data of
// a transaction.
type TXer interface {
	io.Serializable
}
