// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package global

import "errors"

var (
	UserNotFound  = errors.New("用户不存在")
	WrongPassWord = errors.New("密码错误")
)
