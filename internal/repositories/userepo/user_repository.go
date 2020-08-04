// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package userepo

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/datastore/entity"
	"github.com/glepnir/gin-web/internal/repositories"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	conn *gorm.DB
}

func NewUserRepository(c *gin.Context) repositories.UserRepository {
	if conn, ok := c.Value("db").(*gorm.DB); ok {
		return &userRepo{conn}
	} else {
		return &userRepo{}
	}
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
