// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/handlers"
)

func RoleRoute(r *gin.RouterGroup) {
	rolehandler := handlers.NewRoleHandler()
	roleg := r.Group("/roles")
	{
		roleg.POST("", rolehandler.AddRole)
	}
}
