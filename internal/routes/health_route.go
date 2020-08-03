// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthRouter(g *gin.RouterGroup) {
	{
		g.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "Health Ok",
				"data": gin.H{
					"database": "ok",
				},
				"errors": nil,
			})
		})
	}
}
