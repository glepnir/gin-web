// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package userepo

import (
	"github.com/glepnir/gin-web/internal/datastore/entity"
	"github.com/glepnir/gin-web/internal/repositories"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	conn *gorm.DB
}

func NewUserRepository(gdb *gorm.DB) repositories.UserRepository {
	return &userRepo{conn: gdb}
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
