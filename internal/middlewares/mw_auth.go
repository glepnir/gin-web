// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/global"
	"github.com/glepnir/gin-web/pkg/ginresp"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("edu-token")
		if err != nil {
			ginresp.UnAuthorized(c, "无权访问请登录后访问", nil, nil)
			c.Redirect(302, "/")
			c.Abort()
			return
		}
		if token == "" {
			ginresp.UnAuthorized(c, "无权访问请登录后访问", nil, nil)
			c.Redirect(302, "/")
			c.Abort()
			return
		}
		auth := global.NewAuth()
		userid, err := auth.ParseUserID(token)
		if err != nil {
			c.Redirect(302, "/")
			ginresp.UnAuthorized(c, "认证失败请重新登陆", nil, err)
			return
		}
		c.Set("USERID", userid)
		c.Next()
	}
}
