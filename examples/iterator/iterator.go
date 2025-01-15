package iteratorcontract

import (
	"github.com/ixje/neo-go-legacy/pkg/interop/iterator"
	"github.com/ixje/neo-go-legacy/pkg/interop/runtime"
	"github.com/ixje/neo-go-legacy/pkg/interop/storage"
)

// Main is Main(), really.
func Main() bool {
	iter := storage.Find(storage.GetContext(), []byte("foo"))
	values := iterator.Values(iter)
	keys := iterator.Keys(iter)

	runtime.Notify("found storage values", values)
	runtime.Notify("found storage keys", keys)

	return true
}
