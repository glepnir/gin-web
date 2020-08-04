// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/handlers"
	"github.com/glepnir/gin-web/internal/repositories/userepo"
	"github.com/glepnir/gin-web/internal/services/userservice"
)

func UserRoute(g *gin.RouterGroup, c *gin.Context) {
	userg := g.Group("/user")
	userRepository := userepo.NewUserRepository(c)
	userService := userservice.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	{
		userg.POST("/create", userHandler.Create)
	}
}
