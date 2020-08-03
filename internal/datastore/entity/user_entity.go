// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

// User struct for user entity
type User struct {
	Base
	Username string `gorm:"type:varchar(50);unique_index" json:"username"`
	Email    string `gorm:"type:varchar(100);unique" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}
