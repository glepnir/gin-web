// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package userservice

import (
	"time"

	"github.com/glepnir/gin-web/internal/repositories"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/services"
	"github.com/glepnir/gin-web/internal/storage/entity"
	"github.com/glepnir/gin-web/pkg/hash"
)

type userServ struct {
	userRepository repositories.UserRepository
}

func NewUserService(u repositories.UserRepository) services.UserServices {
	return &userServ{userRepository: u}
}

var _UserServ = (*services.UserServices)(nil)

func (u *userServ) CreateUser(userschema schema.CreateUserSchema) (error, bool) {
	_, exist := u.userRepository.UserExist(userschema.Phone)
	if exist {
		userschema.PassWord = hash.HashAndSalt([]byte(userschema.PassWord))
		localtime, _ := time.ParseInLocation("2006-01-02 15:04:05", userschema.ExpireTime, time.Local)
		user := entity.User{
			UserName:       userschema.UserName,
			PassWord:       userschema.PassWord,
			Phone:          userschema.Phone,
			CompanyName:    userschema.CompanyName,
			CompanyAddress: userschema.CompanyAddress,
			Status:         userschema.Status,
			ExpireTime:     localtime,
		}
		_, err := u.userRepository.CreateUser(user)
		if err != nil {
			return err, false
		}
		return nil, true
	} else {
		return nil, false
	}
}

func (u *userServ) UpdateUser(id string, updateuser schema.UpdateUserSchema) error {
	localtime, _ := time.ParseInLocation("2006-01-02 15:04:05", updateuser.ExpireTime, time.Local)
	user := entity.User{
		UserName:       updateuser.UserName,
		Phone:          updateuser.Phone,
		CompanyName:    updateuser.CompanyName,
		CompanyAddress: updateuser.CompanyAddress,
		Status:         updateuser.Status,
		ExpireTime:     localtime,
	}
	err := u.userRepository.UpdateUser(id, user)
	if err != nil {
		return err
	}
	return nil
}
