package domain

type UserState int

const (
	Idle UserState = iota
	Registering
)
