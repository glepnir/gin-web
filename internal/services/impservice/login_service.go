// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impservice

import (
	"os"

	"github.com/glepnir/gin-web/internal/global"
	"github.com/glepnir/gin-web/internal/repositories"
	"github.com/glepnir/gin-web/internal/schema"
	"github.com/glepnir/gin-web/internal/services"
	"github.com/glepnir/gin-web/pkg/auth/jwtauth"
	"github.com/glepnir/gin-web/pkg/hash"
)

type loginServ struct {
	userReader repositories.UserReader
}

func NewLoginServ(u repositories.UserReader) services.LoginServices {
	return &loginServ{u}
}

func (l *loginServ) Login(login schema.LoginSchema) (schema.LoginResultSchema, error) {
	user, exist := l.userReader.GetUserByPhone(login.Phone)
	if exist {
		pass := hash.HashCompare([]byte(user.PassWord), []byte(login.PassWord))
		if pass {
			withexpired := jwtauth.WithExpired(7200)
			pwd, _ := os.Getwd()
			withprivate := jwtauth.WithPrivateKey(pwd + "/configs/app.rsa")
			withpublic := jwtauth.WithPubKeyPath(pwd + "/configs/app.rsa.pub")
			auth := jwtauth.NewJwtAuth(withexpired, withprivate, withpublic)

			token, _ := auth.GenerateToken(user.ID.String())
			return schema.LoginResultSchema{AccessToken: token.AccessToken}, nil
		} else {
			return schema.LoginResultSchema{}, global.WrongPassWord
		}
	} else {
		return schema.LoginResultSchema{}, global.UserNotFound
	}
}
