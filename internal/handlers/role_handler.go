// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/pkg/ginresp"
)

type RoleHandler struct{}

func NewRoleHandler() *RoleHandler {
	return new(RoleHandler)
}

func (r *RoleHandler) AddRole(c *gin.Context) {
	var role schema.RoleSchema
	c.ShouldBindBodyWith(&role, binding.JSON)
	result, err := role.AddRole()
	if err != nil {
		ginresp.InternalError(c, "添加失败", nil, err)
		return
	} else {
		if result {
			ginresp.Ok(c, "添加成功", nil, nil)
		} else {
			ginresp.InternalError(c, "添加失败", nil, err)
		}
	}
}
