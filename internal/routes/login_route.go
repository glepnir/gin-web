// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/handlers"
	"github.com/glepnir/gin-web/internal/repositories/userepo"
	"github.com/glepnir/gin-web/internal/services/userservice"
	"github.com/glepnir/gin-web/internal/storage"
)

func LoginRoute(r *gin.RouterGroup) {
	conn := &storage.DB{}
	userRepository := userepo.NewUserRepository(conn.Get())
	userService := userservice.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	r.POST("/login", userHandler.Login)
}
