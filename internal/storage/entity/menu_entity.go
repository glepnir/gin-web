// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

type Menu struct {
	Base
	MenuName string `gorm:"column:name;size:50;default:'';not null;"`
	Path     string `gorm:"column:path;size:50;default:'';not null;"`
	Type     string `gorm:"column:type;size:50;default:'';not null;"`
	Method   string `gorm:"column:method;size:50;default:'';not null;"`
}
