package latch

const (
	Version   = "1.6"
	LatchHost = "https://latch.elevenpaths.com"
)

// Parameter names
const (
	ParentIdParameter      = "parentId"
	NameParameter          = "name"
	TwoFactorParameter     = "two_factor"
	LockOnRequestParameter = "lock_on_request"
	ContactPhone           = "contactPhone"
	ContactEmail           = "contactEmail"
)

type Credentials struct {
	Id     string
	Secret string
}

// Holds allowed values for two_factor parameter
var TwoFactor = struct {
	Mandatory, OptIn, Disabled string
}{"MANDATORY", "OPT_IN", "DISABLED"}

// Holds allowed values for lock_on_request parameter
var LockOnRequest = struct {
	Mandatory, OptIn, Disabled string
}{"MANDATORY", "OPT_IN", "DISABLED"}
