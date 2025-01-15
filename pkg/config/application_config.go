package config

import (
	"time"

	"github.com/ixje/neo-go-legacy/pkg/core/storage"
	"github.com/ixje/neo-go-legacy/pkg/network/metrics"
	"github.com/ixje/neo-go-legacy/pkg/rpc"
	"github.com/ixje/neo-go-legacy/pkg/wallet"
)

// ApplicationConfiguration config specific to the node.
type ApplicationConfiguration struct {
	Address           string                  `yaml:"Address"`
	AttemptConnPeers  int                     `yaml:"AttemptConnPeers"`
	DBConfiguration   storage.DBConfiguration `yaml:"DBConfiguration"`
	DialTimeout       time.Duration           `yaml:"DialTimeout"`
	LogPath           string                  `yaml:"LogPath"`
	MaxPeers          int                     `yaml:"MaxPeers"`
	MinPeers          int                     `yaml:"MinPeers"`
	NodePort          uint16                  `yaml:"NodePort"`
	PingInterval      time.Duration           `yaml:"PingInterval"`
	PingTimeout       time.Duration           `yaml:"PingTimeout"`
	Pprof             metrics.Config          `yaml:"Pprof"`
	Prometheus        metrics.Config          `yaml:"Prometheus"`
	ProtoTickInterval time.Duration           `yaml:"ProtoTickInterval"`
	Relay             bool                    `yaml:"Relay"`
	RPC               rpc.Config              `yaml:"RPC"`
	UnlockWallet      wallet.Config           `yaml:"UnlockWallet"`
}
