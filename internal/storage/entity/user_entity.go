// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

// User struct for user entity
type User struct {
	Base
	UserName string `gorm:"column:name;size:64;index;default:'';not null;" json:"username"`
	PassWord string `gorm:"column:password;size:60;default:'';not null;" json:"password"`
	Phone    string `gorm:"column:phone;size:20;index;default:'';not null;" json:"phone"`
}
