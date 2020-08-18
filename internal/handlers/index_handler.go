// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/repositories"
)

type IndexHandler struct {
	userRepo repositories.UserReader
}

func NewIndexHandler(userrepo repositories.UserReader) *IndexHandler {
	return &IndexHandler{userrepo}
}

func (i *IndexHandler) IndexHandler(c *gin.Context) {
	user, exist := i.userRepo.GetUserByID(c.MustGet("USERID").(string))
	if !exist && c.Request.Method == http.MethodGet {
		fmt.Println(user.UserName)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": user.UserName,
			"userid":   user.ID.String(),
		})
	}
}
