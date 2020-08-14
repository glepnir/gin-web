// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/services"
	"github.com/glepnir/gin-web/pkg/auth/jwtauth"
	"github.com/glepnir/gin-web/pkg/ginresp"
	"github.com/glepnir/gin-web/pkg/hash"
	"github.com/glepnir/gin-web/pkg/validator"
)

type UserHandler struct {
	userService services.UserServices
}

func NewUserHandler(u services.UserServices) *UserHandler {
	return &UserHandler{userService: u}
}

func (u *UserHandler) Login(c *gin.Context) {
	var login schema.LoginSchema
	err := c.ShouldBindBodyWith(&login, binding.JSON)
	if err != nil {
		ginresp.BadRequest(c, "请求错误", nil, err)
	}

	user, exist := u.userService.GetUserByPhone(login.Phone)
	if exist {
		pass := hash.HashCompare([]byte(user.PassWord), []byte(login.PassWord))
		if pass {
			withexpired := jwtauth.WithExpired(7200)
			pwd, _ := os.Getwd()

			withprivate := jwtauth.WithPrivateKey(pwd + "/configs/app.rsa")
			withpublic := jwtauth.WithPubKeyPath(pwd + "/configs/app.rsa.pub")
			auth := jwtauth.NewJwtAuth(withexpired, withprivate, withpublic)
			token, _ := auth.GenerateToken(user.ID.String())

			ginresp.Ok(c, "登陆成功", token.AccessToken, nil)
		} else {
			ginresp.Ok(c, "密码不正确", nil, nil)
		}
	} else {
		ginresp.NotFound(c, "用户不存在", nil, nil)
	}
}

func (u *UserHandler) Create(c *gin.Context) {
	var user schema.CreateUserSchema
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	err := validator.Validate(&user)
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
	var user schema.UserSchema
	param := schema.UserID{}
	_ = c.ShouldBindUri(&param)
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	err := validator.Validate(&user)
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

func (u *UserHandler) GetUsers(c *gin.Context) {
	users := u.userService.GetUsers()
	ginresp.Ok(c, "", users, nil)
}

func (u *UserHandler) GetUserById(c *gin.Context) {
	var param schema.UserID
	_ = c.ShouldBindUri(&param)
	user, ok := u.userService.GetUserByID(param.ID)
	if ok {
		ginresp.Ok(c, "", user, nil)
	} else {
		ginresp.NotFound(c, "未查到该用户", nil, nil)
	}
}
