// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package userepo

import (
	"github.com/glepnir/gin-web/internal/repositories"
	"github.com/glepnir/gin-web/internal/storage/entity"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repositories.UserRepository {
	return &userRepo{conn}
}

var _ repositories.UserRepository = (*userRepo)(nil)

// Register is method to create user in db
func (r *userRepo) CreateUser(user entity.User) (entity.User, error) {
	userCreated := entity.User{}
	tx := r.conn.Begin()
	if err := tx.Error; err != nil {
		return userCreated, err
	}
	if err := tx.Create(&user).Scan(&userCreated).Error; err != nil {
		tx.Rollback()
		return userCreated, err
	}
	return userCreated, tx.Commit().Error
}

func (r *userRepo) UserExist(phone string) (entity.User, bool) {
	var user entity.User
	exist := r.conn.Select("phone").Where("phone = ?", phone).First(&user).RecordNotFound()
	return user, exist
}

func (r *userRepo) UpdateUser(phone string, update entity.User) error {
	tx := r.conn.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Where("phone = ?", phone).Update(update).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
