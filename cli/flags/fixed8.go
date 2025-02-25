package flags

import (
	"flag"
	"strings"

	"github.com/ixje/neo-go-legacy/pkg/util"
	"github.com/urfave/cli"
)

// Fixed8 is a wrapper for Uint160 with flag.Value methods.
type Fixed8 struct {
	Value util.Fixed8
}

// Fixed8Flag is a flag with type string.
type Fixed8Flag struct {
	Name  string
	Usage string
	Value Fixed8
}

var (
	_ flag.Value = (*Fixed8)(nil)
	_ cli.Flag   = Fixed8Flag{}
)

// String implements fmt.Stringer interface.
func (a Fixed8) String() string {
	return a.Value.String()
}

// Set implements flag.Value interface.
func (a *Fixed8) Set(s string) error {
	f, err := util.Fixed8FromString(s)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	a.Value = f
	return nil
}

// Fixed8 casts address to util.Fixed8.
func (a *Fixed8) Fixed8() util.Fixed8 {
	return a.Value
}

// String returns a readable representation of this value
// (for usage defaults).
func (f Fixed8Flag) String() string {
	var names []string
	eachName(f.Name, func(name string) {
		names = append(names, getNameHelp(name))
	})

	return strings.Join(names, ", ") + "\t" + f.Usage
}

// GetName returns the name of the flag.
func (f Fixed8Flag) GetName() string {
	return f.Name
}

// Apply populates the flag given the flag set and environment
// Ignores errors.
func (f Fixed8Flag) Apply(set *flag.FlagSet) {
	eachName(f.Name, func(name string) {
		set.Var(&f.Value, name, f.Usage)
	})
}

// Fixed8FromContext returns parsed util.Fixed8 value provided flag name.
func Fixed8FromContext(ctx *cli.Context, name string) util.Fixed8 {
	return ctx.Generic(name).(*Fixed8).Value
}
