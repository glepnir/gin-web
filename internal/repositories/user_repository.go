// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repositories

import (
	"github.com/glepnir/gin-web/internal/datastore/entity"
	"github.com/jinzhu/gorm"
)

type UserWriter interface {
	CreateUser(user entity.User) (entity.User, error)
}

type UserReader interface {
}

type UserRepository interface {
	UserWriter
	UserReader
}

func NewUserRepository(gdb *gorm.DB) UserRepository {
	return &User{conn: gdb}
}

type User struct {
	conn *gorm.DB
}

var _ UserRepository = (*User)(nil)

// Register is method to create user in db
func (r *User) CreateUser(user entity.User) (entity.User, error) {
	userCreated := entity.User{}
	if err := r.conn.Create(&user).Scan(&userCreated).Error; err != nil {
		return userCreated, err
	}

	return userCreated, nil
}
