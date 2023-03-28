package api

type RidesService struct{}

func NewRidesService() Service {
	return &RidesService{}
}
