// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/datastore"
	"github.com/glepnir/gin-web/internal/repositories"
	"github.com/glepnir/gin-web/internal/services"
)

func UserRoute(g *gin.RouterGroup) {
	userg := g.Group("/user")
	conn := &datastore.DB{}
	userRepository := repositories.NewUserRepository(conn.Get())
	userService := services.NewUserService(userRepository)
	{
		userg.POST("/create", userControler)
	}
}
