package gorm

import (
	"errors"
	"log"

	"github.com/ValerySidorin/whisper/internal/infrastructure/config"
	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/dto/storage"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type GormStorage struct {
	gormDB *gorm.DB
}

func Register(cfg *config.Configuration) (*GormStorage, error) {
	opts, err := NewGormOptions(cfg.Storage.Options)
	if err != nil {
		return nil, err
	}
	log.Println(opts.Dsn)
	var db = &gorm.DB{}
	switch {
	case opts.Driver == "postgres":
		db, err = gorm.Open(postgres.Open(opts.Dsn))
		if err != nil {
			return nil, err
		}
	case opts.Driver == "sqlserver":
		db, err = gorm.Open(sqlserver.Open(opts.Dsn))
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid storage driver")
	}
	db.AutoMigrate(&storage.User{})
	return &GormStorage{
		gormDB: db,
	}, nil
}

func (gs *GormStorage) AddUser(u *storage.User) (*storage.User, error) {
	result := gs.gormDB.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (gs *GormStorage) GetUserByMessenger(vcsType dto.VCSHostingType, messengerType dto.MessengerType, messengerUserID int64) (*storage.User, error) {
	u := storage.User{}
	result := gs.gormDB.Where("vcs_hosting_type = ? and messenger_type = ? and messenger_user_id = ?", vcsType, messengerType, messengerUserID).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}

func (gs *GormStorage) GetUserByVCSHosting(vcsType dto.VCSHostingType, messengerType dto.MessengerType, vcsHostingUserID int64) (*storage.User, error) {
	u := storage.User{}
	result := gs.gormDB.Where("vcs_hosting_type = ? and messenger_type = ? and vcs_hosting_user_id = ?", vcsType, messengerType, vcsHostingUserID).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}

func (gs *GormStorage) UpdateUser(u *storage.User) (*storage.User, error) {
	result := gs.gormDB.Save(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
