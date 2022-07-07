package gorm

import "errors"

type GormOptions struct {
	Driver string
	Dsn    string
}

func NewGormOptions(opts map[string]interface{}) (*GormOptions, error) {
	driver, ok := opts["driver"]
	if !ok {
		return nil, errors.New("gorm driver is not present")
	}
	dsn, ok := opts["dsn"]
	if !ok {
		return nil, errors.New("dsn string is not present")
	}
	return &GormOptions{
		Driver: driver.(string),
		Dsn:    dsn.(string),
	}, nil
}
