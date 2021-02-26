package repository

import "github.com/DandDevy/glofox/internal"

// BookingRepository is the repository providing accessing to persisted bookings.
type BookingRepository struct {
	db *DB
}

// NewBookingRepository provides a new BookingRepository.
func NewBookingRepository(db *DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) Add(booking internal.Booking) (*internal.Booking, error) {
	r.db.bookings = append(r.db.bookings, &booking)
	return &booking, nil
}
