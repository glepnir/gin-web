// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services

import (
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/storage/entity"
)

type UserServiceWriter interface {
	CreateUser(user schema.CreateUserSchema) (error, bool)
	UpdateUser(id string, user schema.UserSchema) error
}

type UserServiceReader interface {
	GetUsers(currentpage int) (map[string]interface{}, error)
	GetUserByID(id string) (entity.User, bool)
	GetUserByPhone(phone string) (entity.User, bool)
}

type UserServices interface {
	UserServiceWriter
	UserServiceReader
}
