// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impservice

import (
	"github.com/glepnir/gin-web/internal/repositories"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/services"
	"github.com/glepnir/gin-web/internal/storage/entity"
	"github.com/jinzhu/copier"
)

type studentServ struct {
	studentRepo repositories.StudentRepository
}

var _ services.StudentServices = (*studentServ)(nil)

func NewStudentServ(studentrepo repositories.StudentRepository) *studentServ {
	return &studentServ{studentrepo}
}

func (s *studentServ) CreateStudent(student schema.CreateStudentSchema) error {
	var studententity entity.Student
	copier.Copy(&studententity, &student)
	return s.studentRepo.CreateStudent(studententity)
}
