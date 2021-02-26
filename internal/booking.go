package internal

import "errors"

// BookingRepository interface provides access to persisted bookings.
type BookingRepository interface {
	// Add adds a new booking.
	Add(booking Booking) (*Booking, error)
}

// BookingService implements BookingRepository and provides service to bookings.
type BookingService struct {
	repo BookingRepository
	services *ServiceProvider
}

// Add adds a new booking.
func (s *BookingService) Add(booking Booking) (*Booking, error) {
	if len(booking.Name) < 1 {
		return nil, errors.New("no name")
	}

	_, err := s.services.Class().Get(booking.Date)
	if err != nil {
		return nil, err
	}

	newBooking, err := s.repo.Add(booking)
	if err != nil {
		return nil, err
	}

	return newBooking, nil
}

// Booking represents a booking in internal Glofox.
type Booking struct {
	Name string `json:"name"`
	Date string `json:"date"`
}
