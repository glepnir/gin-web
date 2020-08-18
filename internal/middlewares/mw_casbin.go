// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/repositories/imprepository"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/storage"
	"github.com/glepnir/gin-web/pkg/ginresp"
)

func CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, ok := c.Keys["USERID"].(string)
		if ok {
			conn := &storage.DB{}
			userRepo := imprepository.NewUserRepository(conn.Get())
			rolename, exist := userRepo.GetUserRoleName(userid)
			fmt.Println(rolename)
			if exist {
				path := c.Request.URL.Path
				method := c.Request.Method
				e := schema.LoadCasbin()
				res, err := e.Enforce(rolename, path, method)
				if err != nil {
					log.Fatal(err)
					ginresp.InternalError(c, "异常错误", nil, err)
					c.Abort()
					return
				}
				if res {
					c.Next()
				}
			} else {
				ginresp.Forbidden(c, "无权限访问", nil, nil)
				c.Abort()
				return
			}
		} else {
			ginresp.InternalError(c, "请登录后访问", nil, nil)
		}
	}
}
