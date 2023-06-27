package service

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s *Service) GetText() string {
	return "aj git text"
}
