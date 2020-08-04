// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/glepnir/gin-web/internal/datastore/entity"
	"github.com/glepnir/gin-web/internal/services"
)

type UserHandler struct {
	userService services.UserServices
}

func NewUserHandler(u services.UserServices) *UserHandler {
	return &UserHandler{userService: u}
}

func (u *UserHandler) Create(c *gin.Context) {
	var user entity.User
	_ = c.ShouldBindBodyWith(&user, binding.JSON)
	_, err := u.userService.CreateUser(user)
	if err != nil {
		fmt.Fprintf(c.Writer, "success")
	}
}
