package events

// ServiceInterface service interface
type ServiceInterface interface {
	Append(event Event) error
}

// Service service
type Service struct {
	store StoreInterface
}

// NewService create service
func NewService() ServiceInterface {
	return &Service{
		store: NewArrayStore(),
	}
}

// Append append event
func (service *Service) Append(event Event) error {
	return nil
}
