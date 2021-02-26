package repository

import (
	"errors"
	"github.com/DandDevy/glofox/internal"
	"time"
)

// ClassRepository is the repository providing accessing to persisted classes.
type ClassRepository struct {
	db *DB
}

// NewClassRepository provides a new ClassRepository.
func NewClassRepository(db *DB) *ClassRepository {
	return &ClassRepository{db: db}
}

// Add creates a new Class.
func (r *ClassRepository) Add(class internal.Class) (*internal.Class, error) {
	r.db.classes = append(r.db.classes, &class)
	return &class, nil
}

// List classes.
func (r *ClassRepository) List() []*internal.Class {
	classes := r.db.classes
	if classes == nil {
		return []*internal.Class{}
	}

	return classes
}

// Get class for date.
func (r *ClassRepository) Get(date string) (*internal.Class, error) {
	classes := r.db.classes
	if classes == nil {
		return nil, errors.New("not an available date")
	}

	location, err := time.LoadLocation(internal.TimeZone)
	if err != nil {
		return  nil, err
	}
	bookingDate, err := time.ParseInLocation(date, internal.TimeFormat, location)
	if err != nil {
		return nil, err
	}

	for _, class := range classes {
		startDate, err := time.ParseInLocation(class.StartDate, internal.TimeFormat, location)
		if err != nil {
			return nil, err
		}
		endDate, err := time.ParseInLocation(class.EndDate, internal.TimeFormat, location)
		if err != nil {
			return nil, err
		}
		if bookingDate.After(startDate) || bookingDate.Equal(startDate) {
			if bookingDate.Before(endDate) || bookingDate.Equal(endDate) {
				return class, nil
			}
		}
	}

	return nil, errors.New("not an available date")
}
