package repo

import (
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) Select() (*User, error) {
	var user User

	res := r.db.First(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (r *repo) Insert(id string) error {
	var user1 User

	res := r.db.First(&user1)
	if res.Error != nil {
		return res.Error
	}

	user2 := &User{
		ID: id,
	}

	res = r.db.Create(&user2)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
