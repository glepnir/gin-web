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

func StudentRoutes(g *gin.RouterGroup) {
	conn := &storage.DB{}
	studentRepo := imprepository.NewStudentRepo(conn.Get())
	studentServ := impservice.NewStudentServ(studentRepo)
	studentHandler := handlers.NewStudentHandler(studentServ)
	studentg := g.Group("/student")
	{
		studentg.POST("", studentHandler.CreateStudent)
	}

}
