package contract

const (
	STOPING int = iota
	STARTING
	WAIT_TRYING
)

type IProcess interface {
	Stop()
	Start()
	CanTryStart() bool
	Alive() bool
	Status() int
}
