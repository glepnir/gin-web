// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repositories

import (
	"github.com/glepnir/gin-web/internal/storage/entity"
)

type UserWriter interface {
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(id string, update entity.User) error
}

type UserReader interface {
	UserExist(name string) (entity.User, bool)
	GetUsers(currentpage int) (map[string]interface{}, error)
	GetUserByID(id string) (entity.User, bool)
	GetUserByPhone(phone string) (entity.User, bool)
	GetUserRoleName(userid string) (string, bool)
}

type UserRepository interface {
	UserWriter
	UserReader
}
