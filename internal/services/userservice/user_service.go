// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package userservice

import (
	"github.com/glepnir/gin-web/internal/repositories"
	"github.com/glepnir/gin-web/internal/services"
	"github.com/glepnir/gin-web/internal/storage/entity"
)

type userServ struct {
	userRepository repositories.UserRepository
}

func NewUserService(u repositories.UserRepository) services.UserServices {
	return &userServ{userRepository: u}
}

var _UserServ = (*services.UserServices)(nil)

func (u *userServ) CreateUser(user entity.User) error {
	_, err := u.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
