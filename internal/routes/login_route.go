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

func LoginRoute(r *gin.RouterGroup) {
	conn := &storage.DB{}
	userRepo := imprepository.NewUserRepository(conn.Get())
	loginServ := impservice.NewLoginServ(userRepo)
	loginHandler := handlers.NewLoginHandler(loginServ)
	r.POST("/login", loginHandler.Login)
}
