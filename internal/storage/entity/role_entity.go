// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

type Role struct {
	Base
	Name     string `gorm:"column:name;size:100;index;default:'';not null;"`
	Sequence int    `gorm:"column:sequence;index;default:0;not null;"`
	Status   int    `gorm:"column:status;index;default:0;not null;"`
	Creator  string `gorm:"column:creator;size:36;"`
}
