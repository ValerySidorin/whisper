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
	dr, ok := driver.(string)
	if !ok {
		return nil, errors.New("driver string type assertion failed")
	}
	d, ok := dsn.(string)
	if !ok {
		return nil, errors.New("dsn string type assertion failed")
	}
	return &gormOptions{
		driver: dr,
		dsn:    d,
	}, nil
}
