package state

import (
	"testing"

	"github.com/ixje/neo-go-legacy/pkg/internal/testserdes"
)

func TestEncodeDecodeStorageItem(t *testing.T) {
	storageItem := &StorageItem{
		Value:   []byte{},
		IsConst: false,
	}

	testserdes.EncodeDecodeBinary(t, storageItem, new(StorageItem))
}
