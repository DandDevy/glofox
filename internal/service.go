package internal

import "github.com/DandDevy/glofox/internal/repository"

// ServiceProvider provides the different services available for Glofox
type ServiceProvider struct {
	repo *repository.RepositoryProvider
}

// NewServiceProvider returns a new ServiceProvider
func NewServiceProvider(repo *repository.RepositoryProvider) *ServiceProvider {
	return &ServiceProvider{repo: repo}
}
