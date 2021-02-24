package repository

// DB represents a database inside Glofox
type DB struct {
	tables []*table
}

// table represents a database table inside Glofox
type table struct {
	row []string
}
