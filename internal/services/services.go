// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services

import (
	"github.com/glepnir/gin-web/internal/datastore/entity"
)

type UserServiceWriter interface {
	CreateUser(user entity.User) (bool, error)
}

type UserServices interface {
	UserServiceWriter
}
