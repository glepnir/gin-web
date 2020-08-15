// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/handlers"
	"github.com/glepnir/gin-web/internal/repositories/imprepository"
	"github.com/glepnir/gin-web/internal/services/impservice"
	"github.com/glepnir/gin-web/internal/storage"
)

func UserRoute(g *gin.RouterGroup) {
	userg := g.Group("/users")
	conn := &storage.DB{}
	userRepository := imprepository.NewUserRepository(conn.Get())
	userService := impservice.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	{
		userg.GET("", userHandler.GetUsers)
		userg.GET(":id", userHandler.GetUserById)
		userg.POST("", userHandler.Create)
		userg.PUT(":id", userHandler.Update)
	}
}
