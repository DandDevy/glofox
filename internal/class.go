package internal

import (
	"errors"
	"time"
)

// ClassRepository interface provides access to persisted classes.
type ClassRepository interface {
	// Add adds a new class.
	Add(class Class) (*Class, error)
	// List classes.
	List() []*Class
	// Get class for date.
	Get(date string) (*Class, error)
}

// ClassService implements ClassRepository and provides service to classes.
type ClassService struct {
	repo ClassRepository
	services *ServiceProvider
}

// Add adds a new class.
func (s *ClassService) Add(class Class) (*Class, error) {
	if class.Capacity < 1 {
		return nil, errors.New("too little capacity")
	}
	if len(class.Name) < 1 {
		return nil, errors.New("no name")
	}

	err := validateDate(class.StartDate)
	if err != nil {
		return nil, err
	}

	err = validateDate(class.EndDate)
	if err != nil {
		return nil, err
	}

	newClass, err := s.repo.Add(class)
	if err != nil {
		return nil, err
	}

	return newClass, nil
}

// List classes.
func (s *ClassService) List() []*Class {
	return s.repo.List()
}

// Get class for date.
func (s *ClassService) Get(date string) (*Class, error) {
	err := validateDate(date)
	if err != nil {
		return nil, err
	}

	return s.repo.Get(date)
}

const TimeZone = "UTC"
const TimeFormat = "2020-Jul-09"

func validateDate(date string)  error {
	location, err := time.LoadLocation(TimeZone)
	if err != nil {
		return  err
	}
	_, err = time.ParseInLocation(date, TimeFormat, location)
	if err != nil {
		return  err
	}
	return nil
}

// Class represents a class in internal Glofox.
type Class struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  uint `json:"capacity"`
}

