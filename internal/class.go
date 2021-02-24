package internal

import (
	"errors"
)

// ClassRepository interface provides access to persisted classes.
type ClassRepository interface {
	// Add adds a new class.
	Add(class Class) (*Class, error)
	// List classes.
	List() []*Class
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

// Class represents a class in internal Glofox
type Class struct {
	Name      string
	StartDate string
	EndDate   string
	Capacity  uint
}

