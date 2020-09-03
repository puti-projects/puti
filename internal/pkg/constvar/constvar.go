// constvar const definition
// the different compare with config const and constvar is these constvar would not change (fixed configuration)

package constvar

const (
	// DefaultLimit default limit
	DefaultLimit = 20
)

const (
	// SecondsPerMinute seconds per minute
	SecondsPerMinute = 60
	// SecondsPerHour seconds per hour
	SecondsPerHour = SecondsPerMinute * 60
	// SecondsPerDay seconds per day
	SecondsPerDay = SecondsPerHour * 24
)
