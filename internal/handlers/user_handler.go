// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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
	fmt.Println(user)
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
	currentPage, _ := strconv.Atoi(c.Query("page_index"))
	if currentPage == 0 {
		currentPage = 1
	}
	models, err := u.userService.GetUsers(currentPage)
	if err != nil {
		ginresp.InternalError(c, "服务器异常", nil, err)
	}
	ginresp.OkWithCount(c, "获取数据成功", models["list"], models["totalItems"].(int), nil)
}

func (u *UserHandler) GetUserById(c *gin.Context) {
	var param schema.UserID
	_ = c.ShouldBindUri(&param)
	currentuser, ok := u.userService.GetUserByID(param.ID)
	if ok {
		c.HTML(http.StatusOK, "admin-edit.html", gin.H{
			"current_userid":             param.ID,
			"current_username":           currentuser.UserName,
			"current_userphone":          currentuser.Phone,
			"current_usercompany":        currentuser.CompanyName,
			"current_usercompanyaddress": currentuser.CompanyAddress,
			"current_expiretime":         currentuser.ExpireTime.String(),
			"current_userstatus":         currentuser.Status,
		})
	} else {
		ginresp.NotFound(c, "未查到该用户", nil, nil)
	}
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	var param schema.UserID
	_ = c.ShouldBindUri(&param)
	err := u.userService.DeleteUser(param.ID)
	if err != nil {
		ginresp.InternalError(c, "删除失败服务器异常", nil, err)
		return
	}
	ginresp.Ok(c, "删除成功", nil, nil)
}
