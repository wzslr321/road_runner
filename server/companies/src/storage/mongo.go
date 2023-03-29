package storage

// Mongo I don't know what db yet, it is just a mock
type Mongo struct{}

func New() *Mongo {
	return &Mongo{}
}
