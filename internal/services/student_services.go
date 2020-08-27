// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services

import "github.com/glepnir/gin-web/internal/schema"

type StudentServices interface {
	CreateStudent(student schema.CreateStudentSchema) error
}
