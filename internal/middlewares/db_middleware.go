// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/config"
	"github.com/glepnir/gin-web/internal/datastore"
)

func SetDBConnection(storage config.Storage) gin.HandlerFunc {
	db, err := datastore.NewDB(storage)
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
