// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/glepnir/gin-web/internal/global"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/services"
	"github.com/glepnir/gin-web/pkg/ginresp"
	"github.com/glepnir/gin-web/pkg/validator"
)

type LoginHandler struct {
	loginservice services.LoginServices
}

func NewLoginHandler(l services.LoginServices) *LoginHandler {
	return &LoginHandler{l}
}

func (l *LoginHandler) Login(c *gin.Context) {
	var login schema.LoginSchema
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	err := c.ShouldBindWith(&login, binding.JSON)
	if err != nil {
		ginresp.BadRequest(c, "请求错误", nil, err)
		return
	}
	err = validator.Validate(login)
	if err != nil {
		ginresp.BadRequest(c, err.Error(), nil, nil)
		return
	}
	loginresult, err := l.loginservice.Login(login)
	if errors.Is(err, global.UserNotFound) {
		ginresp.NotFound(c, err.Error(), nil, nil)
		return
	} else if errors.Is(err, global.WrongPassWord) {
		ginresp.PassWordWrong(c, err.Error(), nil, nil)
		return
	}
	if time.Now().After(loginresult.ExpireTime) {
		ginresp.OkWithFailed(c, "账号已过期", nil, nil)
		return
	}
	c.Set("USERINFO", loginresult)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "edu-token",
		Value:    loginresult.AccessToken,
		MaxAge:   7200,
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	})
	ginresp.Ok(c, "登陆成功", loginresult, nil)
}
