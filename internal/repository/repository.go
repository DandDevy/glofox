package repository

// RepositoryProvider provides access to Glofox repositories.
type RepositoryProvider struct {
	db *DB
}

// NewRepositoryProvider returns a new RepositoryProvider.
func NewRepositoryProvider(db *DB) *RepositoryProvider {
	return &RepositoryProvider{db: db}
}
