package utils

import (
	"github.com/O1MaGnUmO1/chainlink-common/pkg/services"
)

// StartStopOnce can be embedded in a struct to help implement types.Service.
// Deprecated: use services.StateMachine
type StartStopOnce = services.StateMachine
