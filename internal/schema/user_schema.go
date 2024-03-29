// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schema

type UserID struct {
	ID string `uri:"id"`
}

type CreateUserSchema struct {
	UserName       string `json:"username" validate:"required" label:"用户姓名"`
	PassWord       string `json:"password" validate:"required" label:"密码"`
	Phone          string `json:"phone" validate:"required,mobile" label:"联系电话"`
	CompanyName    string `json:"companyname" validate:"required" label:"公司名称"`
	CompanyAddress string `json:"companyaddress" validate:"required" label:"公司地址"`
	ExpireTime     string `json:"expiretime" validate:"required" label:"到期时间"`
	RoleName       string `json:"rolename" validate:"required" label:"权限"`
}

type UserSchema struct {
	UserID
	UserName       string `json:"username" validate:"required" label:"用户姓名"`
	Phone          string `json:"phone" validate:"required,mobile" label:"联系电话"`
	Status         string `json:"status" validate:"required,max=2,min=1" label:"状态"`
	CompanyName    string `json:"companyname" validate:"required" label:"公司名称"`
	CompanyAddress string `json:"companyaddress" validate:"required" label:"公司地址"`
}

type GetUsersSchema struct {
	ID             string
	UserName       string
	Phone          string
	Status         int
	ExpireTime     string
	CompanyName    string
	CompanyAddress string
}
