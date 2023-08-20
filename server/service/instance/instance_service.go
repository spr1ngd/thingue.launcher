package instance

type instanceService struct{}

var InstanceService = new(instanceService)

func (s *instanceService) UpdatePlayer(id uint) {
}
