package instance

type instanceService struct{}

var Service = new(instanceService)

func (s *instanceService) UpdatePlayer(id uint) {
}
