package flags

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ixje/neo-go-legacy/pkg/encoding/address"
	"github.com/ixje/neo-go-legacy/pkg/util"
	"github.com/urfave/cli"
)

// Address is a wrapper for Uint160 with flag.Value methods.
type Address struct {
	IsSet bool
	Value util.Uint160
}

// AddressFlag is a flag with type string
type AddressFlag struct {
	Name  string
	Usage string
	Value Address
}

var (
	_ flag.Value = (*Address)(nil)
	_ cli.Flag   = AddressFlag{}
)

// String implements fmt.Stringer interface.
func (a Address) String() string {
	return address.Uint160ToString(a.Value)
}

// Set implements flag.Value interface.
func (a *Address) Set(s string) error {
	addr, err := address.StringToUint160(s)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	a.IsSet = true
	a.Value = addr
	return nil
}

// Uint160 casts address to Uint160.
func (a *Address) Uint160() (u util.Uint160) {
	if !a.IsSet {
		// It is a programmer error to call this method without
		// checking if the value was provided.
		panic("address was not set")
	}
	return a.Value
}

// IsSet checks if flag was set to a non-default value.
func (f AddressFlag) IsSet() bool {
	return f.Value.IsSet
}

// String returns a readable representation of this value
// (for usage defaults)
func (f AddressFlag) String() string {
	var names []string
	eachName(f.Name, func(name string) {
		names = append(names, getNameHelp(name))
	})

	return strings.Join(names, ", ") + "\t" + f.Usage
}

func getNameHelp(name string) string {
	if len(name) == 1 {
		return fmt.Sprintf("-%s value", name)
	}
	return fmt.Sprintf("--%s value", name)
}

// GetName returns the name of the flag
func (f AddressFlag) GetName() string {
	return f.Name
}

// Apply populates the flag given the flag set and environment
// Ignores errors
func (f AddressFlag) Apply(set *flag.FlagSet) {
	eachName(f.Name, func(name string) {
		set.Var(&f.Value, name, f.Usage)
	})
}
