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
	if err := r.conn.Create(&user).Scan(&userCreated).Error; err != nil {
		return userCreated, err
	}
	return userCreated, nil
}
