package gorm

import "errors"

type gormOptions struct {
	driver string
	dsn    string
}

func newGormOptions(opts map[string]interface{}) (*gormOptions, error) {
	driver, ok := opts["driver"]
	if !ok {
		return nil, errors.New("driver is not present")
	}
	dsn, ok := opts["dsn"]
	if !ok {
		return nil, errors.New("dsn string is not present")
	}
	return &gormOptions{
		driver: driver.(string),
		dsn:    dsn.(string),
	}, nil
}
