package srv

import (
	"github.com/pborman/uuid"
)

// IDGenerator generates command IDs.
// Function is used to ease mocking.
var IDGenerator = func() string {
	return uuid.New()
}
