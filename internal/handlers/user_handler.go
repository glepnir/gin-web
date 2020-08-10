// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/services"
	"github.com/glepnir/gin-web/pkg/ginresp"
	"github.com/glepnir/gin-web/pkg/validator"
)

type UserHandler struct {
	userService services.UserServices
}

func NewUserHandler(u services.UserServices) *UserHandler {
	return &UserHandler{userService: u}
}

func (u *UserHandler) Create(c *gin.Context) {
	var user schema.CreateUserSchema
	_ = c.ShouldBindBodyWith(&user, binding.JSON)
	fmt.Println(user.ExpireTime)
	v := &validator.CustomValidator{}
	err := v.Validate(&user)
	if err != nil {
		ginresp.BadRequest(c, err.Error(), nil, err)
		return
	}

	err, ok := u.userService.CreateUser(user)
	if ok {
		ginresp.Ok(c, "添加用户成功", nil, nil)
	} else {
		if err != nil {
			ginresp.InternalError(c, "服务器异常添加失败", nil, err)
		} else {
			ginresp.Conflict(c, "用户已存在添加失败", nil, nil)
		}
	}

}

func (u *UserHandler) Update(c *gin.Context) {
	var user schema.UpdateUserSchema
	param := schema.UserID{}
	_ = c.ShouldBindUri(&param)
	_ = c.ShouldBindBodyWith(&user, binding.JSON)
	v := new(validator.CustomValidator)
	err := v.Validate(&user)
	if err != nil {
		ginresp.BadRequest(c, err.Error(), nil, nil)
		return
	}

	err = u.userService.UpdateUser(param.ID, user)
	if err != nil {
		ginresp.InternalError(c, "更新失败", nil, err)
	} else {
		ginresp.Ok(c, "更新成功", nil, nil)
	}
}
