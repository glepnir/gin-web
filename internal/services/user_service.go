// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services

import (
	"github.com/glepnir/gin-web/internal/datastore/entity"
	"github.com/glepnir/gin-web/internal/repositories"
)

type UserService interface {
	CreateUser(user entity.User) (bool, error)
}

func NewUserService(u repositories.UserRepository) UserService {
	return &UserResponse{UserRepository: u}
}

type UserResponse struct {
	UserRepository repositories.UserRepository
}

var _UserService = (*UserResponse)(nil)

func (u *UserResponse) CreateUser(user entity.User) (bool, error) {
	return u.CreateUser(user)
}
