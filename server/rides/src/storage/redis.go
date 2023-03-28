package storage

type Redis struct{}

func New() *Redis {
	return &Redis{}
}
