package synchronizer

import (
	"github.com/0xPolygon/cdk-validium-node/config/types"
)

// Config represents the configuration of the synchronizer
type Config struct {
	// SyncInterval is the delay interval between reading new rollup information
	SyncInterval types.Duration `mapstructure:"SyncInterval"`
	// SyncChunkSize is the number of blocks to sync on each chunk
	SyncChunkSize uint64 `mapstructure:"SyncChunkSize"`
	// TrustedSequencerURL is the rpc url to connect and sync the trusted state
	TrustedSequencerURL string `mapstructure:"TrustedSequencerURL"`
	// SyncOnlyTrusted option whether sync L1 block or not (for external node use)
	SyncOnlyTrusted bool `mapstructure:"SyncOnlyTrusted"`
}
