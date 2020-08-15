// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imprepository

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

func (r *userRepo) UpdateUser(id string, update entity.User) error {
	tx := r.conn.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&update).Where("id = ?", id).Update(update).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *userRepo) GetUsers() []entity.User {
	var users []entity.User
	r.conn.Select(
		"id,username,phone,status,companyname,companyaddress,expiretime,created_at,updated_at").Where("deleted_at is null").First(&users)
	return users
}

func (r *userRepo) GetUserByID(id string) (entity.User, bool) {
	var user entity.User
	exist := r.conn.Select(
		"id,username,phone,status,companyname,companyaddress,expiretime,created_at,updated_at").Where("id = ?", id).First(&user).RecordNotFound()
	return user, exist
}

func (r *userRepo) GetUserByPhone(phone string) (entity.User, bool) {
	var user entity.User
	exist := r.conn.Where("phone = ?", phone).First(&user).RecordNotFound()
	if exist {
		return user, false
	}
	return user, true
}
