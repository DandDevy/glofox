package repository

import "github.com/DandDevy/glofox/internal"

// RepositoryProvider provides access to Glofox repositories.
type RepositoryProvider struct {
	db *DB
}

// NewRepositoryProvider returns a new RepositoryProvider.
func NewRepositoryProvider(db *DB) *RepositoryProvider {
	return &RepositoryProvider{db: db}
}

func (r *RepositoryProvider) Class() internal.ClassRepository {
	return NewClassRepository(r.db)
}

func (r *RepositoryProvider) Booking() internal.BookingRepository {
	return NewBookingRepository(r.db)
}


