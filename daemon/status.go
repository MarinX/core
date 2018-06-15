package daemon

import (
	"github.com/mesg-foundation/core/container"
)

// Status returns the Status of the docker service of the daemon
func Status() (status container.StatusType, err error) {
	return container.ServiceStatus(Namespace()) //TODO: should it be containerStatus?
}
