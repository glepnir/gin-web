// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/services"
	"github.com/glepnir/gin-web/internal/storage/entity"
	"github.com/glepnir/gin-web/pkg/ginresp"
	"github.com/glepnir/gin-web/pkg/hash"
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
	hashpwd := hash.HashAndSalt([]byte(user.PassWord))
	user.PassWord = hashpwd

	err, ok := u.userService.CreateUser(user)
	if ok {
		ginresp.Ok(c, "Create user success", nil, nil)
	} else {
		if err != nil {
			ginresp.InternalError(c, "Create failed", nil, err)
		} else {
			ginresp.Conflict(c, "User Exist Create failed", nil, nil)
		}
	}

}

func (u *UserHandler) Update(c *gin.Context) {
	var user entity.User
	param := schema.UserID{}
	_ = c.ShouldBindUri(&param)
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	err := u.userService.UpdateUser(param.ID, user)
	if err != nil {
		ginresp.InternalError(c, "更新失败", nil, err)
	} else {
		ginresp.Ok(c, "更新成功", nil, nil)
	}
}
