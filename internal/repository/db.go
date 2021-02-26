package repository

import "github.com/DandDevy/glofox/internal"

// DB represents a database inside Glofox
type DB struct {
	classes  []*internal.Class
	bookings []*internal.Booking
}

