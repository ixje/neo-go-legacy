package keys

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/ixje/neo-go-legacy/pkg/internal/keytestcases"
	"github.com/stretchr/testify/assert"
)

func TestPrivateKey(t *testing.T) {
	for _, testCase := range keytestcases.Arr {
		privKey, err := NewPrivateKeyFromHex(testCase.PrivateKey)
		if testCase.Invalid {
			assert.Error(t, err)
			continue
		}

		assert.Nil(t, err)
		address := privKey.Address()
		assert.Equal(t, testCase.Address, address)

		wif := privKey.WIF()
		assert.Equal(t, testCase.Wif, wif)
		pubKey := privKey.PublicKey()
		assert.Equal(t, hex.EncodeToString(pubKey.Bytes()), testCase.PublicKey)
	}
}

func TestPrivateKeyFromWIF(t *testing.T) {
	for _, testCase := range keytestcases.Arr {
		key, err := NewPrivateKeyFromWIF(testCase.Wif)
		if testCase.Invalid {
			assert.Error(t, err)
			continue
		}

		assert.Nil(t, err)
		assert.Equal(t, testCase.PrivateKey, key.String())
	}
}

func TestSigning(t *testing.T) {
	// These were taken from the rfcPage:https://tools.ietf.org/html/rfc6979#page-33
	//   public key: U = xG
	//Ux = 60FED4BA255A9D31C961EB74C6356D68C049B8923B61FA6CE669622E60F29FB6
	//Uy = 7903FE1008B8BC99A41AE9E95628BC64F2F1B20C2D7E9F5177A3C294D4462299
	PrivateKey, _ := NewPrivateKeyFromHex("C9AFA9D845BA75166B5C215767B1D6934E50C3DB36E89B127B8A622B120F6721")

	data := PrivateKey.Sign([]byte("sample"))

	r := "EFD48B2AACB6A8FD1140DD9CD45E81D69D2C877B56AAF991C34D0EA84EAF3716"
	s := "F7CB1C942D657C41D436C7A1B6E29F65F3E900DBB9AFF4064DC4AB2F843ACDA8"
	assert.Equal(t, strings.ToLower(r+s), hex.EncodeToString(data))
}
