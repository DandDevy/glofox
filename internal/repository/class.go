package repository

import (
	"github.com/DandDevy/glofox/internal"
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
