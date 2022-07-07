package dto

type UserState string

const (
	Idle        UserState = "idle"
	Registering UserState = "registering"
)

type MessengerType string

const (
	Telegram MessengerType = "telegram"
)

type VCSHostingType string

const (
	Gitlab VCSHostingType = "gitlab"
)

type StorageType string

const (
	Gorm StorageType = "gorm"
)
