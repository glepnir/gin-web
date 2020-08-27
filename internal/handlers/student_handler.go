// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/services"
	"github.com/glepnir/gin-web/pkg/ginresp"
	"github.com/glepnir/gin-web/pkg/validator"
)

type StudentHandler struct {
	studentServ services.StudentServices
}

func NewStudentHandler(studentServ services.StudentServices) *StudentHandler {
	return &StudentHandler{studentServ}
}

func (s *StudentHandler) CreateStudent(c *gin.Context) {
	var student schema.CreateStudentSchema
	err := c.ShouldBindBodyWith(&student, binding.JSON)
	if err != nil {
		ginresp.BadRequest(c, "请求失败", nil, nil)
	}
	err = validator.Validate(student)
	if err != nil {
		ginresp.BadRequest(c, err.Error(), nil, err)
		return
	}
	err = s.studentServ.CreateStudent(student)
	if err != nil {
		ginresp.Ok(c, "添加学生成功", nil, nil)
	} else {
		ginresp.OkWithFailed(c, "添加学生失败", nil, nil)
	}
}
