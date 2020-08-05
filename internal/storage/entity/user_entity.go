// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

// User struct for user entity
type User struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}