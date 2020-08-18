// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schema

type LoginSchema struct {
	Phone    string `json:"phone" validate:"required,mobile" label:"联系电话"`
	PassWord string `json:"password" validate:"required,min=6,max=12" label:"密码"`
}

type LoginResultSchema struct {
	AccessToken string `json:"access_token"`
	UserName    string `json:"username"`
	Phone       string `json:"phone"`
	CompanyName string `json:"companyname"`
}
