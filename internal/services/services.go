// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services

import (
	"github.com/glepnir/gin-web/internal/schema"
)

type UserServiceWriter interface {
	CreateUser(user schema.CreateUserSchema) (error, bool)
	UpdateUser(id string, user schema.UpdateUserSchema) error
}

type UserServices interface {
	UserServiceWriter
}
