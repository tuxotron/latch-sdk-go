package latch

const (
	Version   = "1.6"
	LatchHost = "https://latch.elevenpaths.com"
)

type Credentials struct {
	Id     string
	Secret string
}
