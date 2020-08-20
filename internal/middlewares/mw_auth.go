// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/global"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("edu-token")
		if err != nil {
			c.Redirect(302, "/")
			c.Abort()
			return
		}
		if token == "" {
			c.Redirect(302, "/")
			c.Abort()
			return
		}
		auth := global.NewAuth()
		userid, err := auth.ParseUserID(token)
		if err != nil {
			c.Redirect(302, "/")
			return
		}
		c.Set("USERID", userid)
		c.Next()
	}
}
