package internal

// RepositoryProvider provides access to different services.
type RepositoryProvider interface {
	Class() ClassRepository
	Booking() BookingRepository

}

// ServiceProvider provides the different services available for Glofox.
type ServiceProvider struct {
	repo RepositoryProvider
}

// NewServiceProvider returns a new ServiceProvider.
func NewServiceProvider(repo RepositoryProvider) *ServiceProvider {
	return &ServiceProvider{repo: repo}
}

// Class returns the a ClassService.
func (s *ServiceProvider) Class() *ClassService  {
	return &ClassService{repo: s.repo.Class(), services: s}
}


func (s *ServiceProvider) Booking() *BookingService {
	return &BookingService{repo: s.repo.Booking(), services: s}
}
