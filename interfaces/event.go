package interfaces

type Event interface {
	SaveToDB() error
	Transform() error
}
