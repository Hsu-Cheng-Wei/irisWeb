package repository

import (
	"irisWeb/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (u *UserRepository) GetAll() []domain.User {
	users := []domain.User{}
	u.Db.Find(&users)

	return users
}

func (r *UserRepository) Insert(u domain.User) ([]byte, error) {

	id, _ := uuid.New().MarshalBinary()

	u.ID = id

	err := r.Db.Create(&u).Error

	return id, err
}

func (r *UserRepository) Delete(id string) error {
	uid, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	b, err := uid.MarshalBinary()

	if err != nil {
		return err
	}

	r.Db.Delete(&domain.User{ID: b})

	return nil
}
