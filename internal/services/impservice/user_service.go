// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impservice

import (
	"strconv"
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
		localtime, _ := time.Parse("2006-01-02", userschema.ExpireTime)
		user := entity.User{
			UserName:       userschema.UserName,
			PassWord:       userschema.PassWord,
			Phone:          userschema.Phone,
			CompanyName:    userschema.CompanyName,
			CompanyAddress: userschema.CompanyAddress,
			Status:         1,
			ExpireTime:     localtime,
			RoleName:       userschema.RoleName,
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

func (u *userServ) UpdateUser(id string, updateuser schema.UserSchema) error {
	status, _ := strconv.Atoi(updateuser.Status)
	user := entity.User{
		UserName:       updateuser.UserName,
		Phone:          updateuser.Phone,
		Status:         status,
		CompanyName:    updateuser.CompanyName,
		CompanyAddress: updateuser.CompanyAddress,
	}
	err := u.userRepository.UpdateUser(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userServ) GetUsers(currentpage, limit int) ([]schema.GetUsersSchema, int, error) {
	users, count, err := u.userRepository.GetUsers(currentpage, limit)
	usersschema := make([]schema.GetUsersSchema, len(users))
	if err == nil {
		for index, user := range users {
			ts := user.ExpireTime.Format("2006-01-02")
			usersschema[index] = schema.GetUsersSchema{
				ID:             user.ID.String(),
				UserName:       user.UserName,
				Phone:          user.Phone,
				Status:         user.Status,
				ExpireTime:     ts,
				CompanyName:    user.CompanyName,
				CompanyAddress: user.CompanyAddress,
			}
		}
	}
	return usersschema, count, err
}

func (u *userServ) GetUserByID(id string) (entity.User, bool) {
	user, exist := u.userRepository.GetUserByID(id)
	if exist {
		return user, false
	}
	return user, true
}

func (u *userServ) GetUserByPhone(phone string) (schema.GetUsersSchema, int, bool) {
	var getuser schema.GetUsersSchema
	user, count, exist := u.userRepository.GetUserByPhone(phone)
	if exist {
		getuser.UserName = user.UserName
		getuser.ID = user.ID.String()
		getuser.Phone = user.Phone
		getuser.CompanyAddress = user.CompanyAddress
		getuser.CompanyName = user.CompanyName
		getuser.ExpireTime = user.ExpireTime.Format("2006-01-02")
		return getuser, count, true
	} else {
		return getuser, count, false
	}
}

func (u *userServ) DeleteUser(id string) error {
	return u.userRepository.DeleteUser(id)
}
