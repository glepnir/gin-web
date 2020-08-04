// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repositories

import "github.com/glepnir/gin-web/internal/datastore/entity"

type UserWriter interface {
	CreateUser(user entity.User) (entity.User, error)
}

type UserReader interface {
}

type UserRepository interface {
	UserWriter
	UserReader
}