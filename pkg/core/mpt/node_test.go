package mpt

import (
	"encoding/json"
	"testing"

	"github.com/ixje/neo-go-legacy/pkg/internal/random"
	"github.com/ixje/neo-go-legacy/pkg/internal/testserdes"
	"github.com/ixje/neo-go-legacy/pkg/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getTestFuncEncode(ok bool, expected, actual Node) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("IO", func(t *testing.T) {
			bs, err := testserdes.EncodeBinary(expected)
			require.NoError(t, err)
			err = testserdes.DecodeBinary(bs, actual)
			if !ok {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, expected.Type(), actual.Type())
			require.Equal(t, expected.Hash(), actual.Hash())
		})
		t.Run("JSON", func(t *testing.T) {
			bs, err := json.Marshal(expected)
			require.NoError(t, err)
			err = json.Unmarshal(bs, actual)
			if !ok {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, expected.Type(), actual.Type())
			require.Equal(t, expected.Hash(), actual.Hash())
		})
	}
}

func TestNode_Serializable(t *testing.T) {
	t.Run("Leaf", func(t *testing.T) {
		t.Run("Good", func(t *testing.T) {
			l := NewLeafNode(random.Bytes(123))
			t.Run("Raw", getTestFuncEncode(true, l, new(LeafNode)))
			t.Run("WithType", getTestFuncEncode(true, &NodeObject{l}, new(NodeObject)))
		})
		t.Run("BigValue", getTestFuncEncode(false,
			NewLeafNode(random.Bytes(MaxValueLength+1)), new(LeafNode)))
	})

	t.Run("Extension", func(t *testing.T) {
		t.Run("Good", func(t *testing.T) {
			e := NewExtensionNode(random.Bytes(42), NewLeafNode(random.Bytes(10)))
			t.Run("Raw", getTestFuncEncode(true, e, new(ExtensionNode)))
			t.Run("WithType", getTestFuncEncode(true, &NodeObject{e}, new(NodeObject)))
		})
		t.Run("BigKey", getTestFuncEncode(false,
			NewExtensionNode(random.Bytes(MaxKeyLength+1), NewLeafNode(random.Bytes(10))), new(ExtensionNode)))
	})

	t.Run("Branch", func(t *testing.T) {
		b := NewBranchNode()
		b.Children[0] = NewLeafNode(random.Bytes(10))
		b.Children[lastChild] = NewHashNode(random.Uint256())
		t.Run("Raw", getTestFuncEncode(true, b, new(BranchNode)))
		t.Run("WithType", getTestFuncEncode(true, &NodeObject{b}, new(NodeObject)))
	})

	t.Run("Hash", func(t *testing.T) {
		t.Run("Good", func(t *testing.T) {
			h := NewHashNode(random.Uint256())
			t.Run("Raw", getTestFuncEncode(true, h, new(HashNode)))
			t.Run("WithType", getTestFuncEncode(true, &NodeObject{h}, new(NodeObject)))
		})
		t.Run("Empty", func(t *testing.T) { // compare nodes, not hashes
			testserdes.EncodeDecodeBinary(t, new(HashNode), new(HashNode))
		})
		t.Run("InvalidSize", func(t *testing.T) {
			buf := io.NewBufBinWriter()
			buf.BinWriter.WriteVarBytes(make([]byte, 13))
			require.Error(t, testserdes.DecodeBinary(buf.Bytes(), new(HashNode)))
		})
	})

	t.Run("Invalid", func(t *testing.T) {
		require.Error(t, testserdes.DecodeBinary([]byte{0xFF}, new(NodeObject)))
	})
}

// https://github.com/neo-project/neo/blob/neox-2.x/neo.UnitTests/UT_MPTTrie.cs#L198
func TestJSONSharp(t *testing.T) {
	tr := NewTrie(nil, false, newTestStore())
	require.NoError(t, tr.Put([]byte{0xac, 0x11}, []byte{0xac, 0x11}))
	require.NoError(t, tr.Put([]byte{0xac, 0x22}, []byte{0xac, 0x22}))
	require.NoError(t, tr.Put([]byte{0xac}, []byte{0xac}))
	require.NoError(t, tr.Delete([]byte{0xac, 0x11}))
	require.NoError(t, tr.Delete([]byte{0xac, 0x22}))

	js, err := tr.root.MarshalJSON()
	require.NoError(t, err)
	require.JSONEq(t, `{"key":"0a0c", "next":{"value":"ac"}}`, string(js))
}

func TestInvalidJSON(t *testing.T) {
	t.Run("InvalidChildrenCount", func(t *testing.T) {
		var cs [childrenCount + 1]Node
		for i := range cs {
			cs[i] = new(HashNode)
		}
		data, err := json.Marshal(cs)
		require.NoError(t, err)

		var n NodeObject
		require.Error(t, json.Unmarshal(data, &n))
	})

	testCases := []struct {
		name string
		data []byte
	}{
		{"WrongFieldCount", []byte(`{"key":"0102", "next": {}, "field": {}}`)},
		{"InvalidField1", []byte(`{"next":{}}`)},
		{"InvalidField2", []byte(`{"key":"0102", "hash":{}}`)},
		{"InvalidKey", []byte(`{"key":"xy", "next":{}}`)},
		{"InvalidNext", []byte(`{"key":"01", "next":[]}`)},
		{"InvalidHash", []byte(`{"hash":"01"}`)},
		{"InvalidValue", []byte(`{"value":1}`)},
		{"InvalidBranch", []byte(`[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]`)},
	}
	for _, tc := range testCases {
		var n NodeObject
		assert.Errorf(t, json.Unmarshal(tc.data, &n), "no error in "+tc.name)
	}
}

// C# interoperability test
// https://github.com/neo-project/neo/blob/neox-2.x/neo.UnitTests/UT_MPTTrie.cs#L135
func TestRootHash(t *testing.T) {
	b := NewBranchNode()
	r := NewExtensionNode([]byte{0x0A, 0x0C}, b)

	v1 := NewLeafNode([]byte{0xAB, 0xCD})
	l1 := NewExtensionNode([]byte{0x01}, v1)
	b.Children[0] = l1

	v2 := NewLeafNode([]byte{0x22, 0x22})
	l2 := NewExtensionNode([]byte{0x09}, v2)
	b.Children[9] = l2

	r1 := NewExtensionNode([]byte{0x0A, 0x0C, 0x00, 0x01}, v1)
	require.Equal(t, "dea3ab46e9461e885ed7091c1e533e0a8030b248d39cbc638962394eaca0fbb3", r1.Hash().StringLE())
	require.Equal(t, "93e8e1ffe2f83dd92fca67330e273bcc811bf64b8f8d9d1b25d5e7366b47d60d", r.Hash().StringLE())
}
