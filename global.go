package latch

const (
	Version   = "1.6"
	LatchHost = "https://latch.elevenpaths.com"
)

type Credentials struct {
	Id     string
	Secret string
}

// Holds allowed values for two_factor parameter
type twoFactor struct {
	Mandatory string
	OptIn     string
	Disabled  string
}

// Holds allowed values for lock_on_request parameter
type lockOnRequest struct {
	Mandatory string
	OptIn     string
	Disabled  string
}

var TwoFactor twoFactor
var LockOnRequest lockOnRequest

func init() {
	TwoFactor = twoFactor{"MANDATORY", "OPT_IN", "DISABLED"}
	LockOnRequest = lockOnRequest{"MANDATORY", "OPT_IN", "DISABLED"}
}
