package chaincfg

import (
	"github.com/mleku/ec/wire"
)

// ConsensusDeploymentStarter determines if a given consensus deployment has
// started. A deployment has started once according to the current "time", the
// deployment is eligible for activation once a perquisite condition has
// passed.
type ConsensusDeploymentStarter interface {
	// HasStarted returns true if the consensus deployment has started.
	HasStarted(*wire.BlockHeader) (bool, error)
}

// ConsensusDeploymentEnder determines if a given consensus deployment has
// ended. A deployment has ended once according got eh current "time", the
// deployment is no longer eligible for activation.
type ConsensusDeploymentEnder interface {
	// HasEnded returns true if the consensus deployment has ended.
	HasEnded(*wire.BlockHeader) (bool, error)
}
