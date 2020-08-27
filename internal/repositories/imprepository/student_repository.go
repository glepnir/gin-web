// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imprepository

import (
	"github.com/glepnir/gin-web/internal/repositories"
	"github.com/glepnir/gin-web/internal/storage"
	"github.com/glepnir/gin-web/internal/storage/entity"
	"github.com/jinzhu/gorm"
)

type studentRepo struct {
	conn *gorm.DB
}

var _ repositories.StudentRepository = (*studentRepo)(nil)

func NewStudentRepo(conn *gorm.DB) *studentRepo {
	return &studentRepo{conn}
}

func (s *studentRepo) CreateStudent(student entity.Student) error {
	return storage.CreateOne(s.conn, student)
}
