// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

import "time"

// User struct for user entity
type User struct {
	Base
	UserName       string    `gorm:"column:username;size:64;index;default:'';not null;" `
	PassWord       string    `gorm:"column:password;size:60;default:'';not null;" json:"-"`
	Phone          string    `gorm:"column:phone;size:20;index;default:'';not null;"`
	Status         int       `gorm:"column:status;index;default:0;not null;"`
	CompanyName    string    `gorm:"column:companyname;index;default:'';" `
	CompanyAddress string    `gorm:"column:companyaddress;default:'';size:255"`
	ExpireTime     time.Time `gorm:"column:expiretime;index"`
	RoleName       string    `gorm:"column:rolename;size:40;default:''"`
}
