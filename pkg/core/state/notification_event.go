package state

import (
	"github.com/ixje/neo-go-legacy/pkg/io"
	"github.com/ixje/neo-go-legacy/pkg/smartcontract"
	"github.com/ixje/neo-go-legacy/pkg/smartcontract/trigger"
	"github.com/ixje/neo-go-legacy/pkg/util"
	"github.com/ixje/neo-go-legacy/pkg/vm"
)

// NotificationEvent is a tuple of scripthash that emitted the StackItem as a
// notification and that item itself.
type NotificationEvent struct {
	ScriptHash util.Uint160
	Item       vm.StackItem
}

// AppExecResult represent the result of the script execution, gathering together
// all resulting notifications, state, stack and other metadata.
type AppExecResult struct {
	TxHash      util.Uint256
	Trigger     trigger.Type
	VMState     string
	GasConsumed util.Fixed8
	Stack       []smartcontract.Parameter
	Events      []NotificationEvent
}

// EncodeBinary implements the Serializable interface.
func (ne *NotificationEvent) EncodeBinary(w *io.BinWriter) {
	w.WriteBytes(ne.ScriptHash[:])
	vm.EncodeBinaryStackItem(ne.Item, w)
}

// DecodeBinary implements the Serializable interface.
func (ne *NotificationEvent) DecodeBinary(r *io.BinReader) {
	r.ReadBytes(ne.ScriptHash[:])
	ne.Item = vm.DecodeBinaryStackItem(r)
}

// EncodeBinary implements the Serializable interface.
func (aer *AppExecResult) EncodeBinary(w *io.BinWriter) {
	w.WriteBytes(aer.TxHash[:])
	w.WriteB(byte(aer.Trigger))
	w.WriteString(aer.VMState)
	aer.GasConsumed.EncodeBinary(w)
	w.WriteArray(aer.Stack)
	w.WriteArray(aer.Events)
}

// DecodeBinary implements the Serializable interface.
func (aer *AppExecResult) DecodeBinary(r *io.BinReader) {
	r.ReadBytes(aer.TxHash[:])
	aer.Trigger = trigger.Type(r.ReadB())
	aer.VMState = r.ReadString()
	aer.GasConsumed.DecodeBinary(r)
	r.ReadArray(&aer.Stack)
	r.ReadArray(&aer.Events)
}
