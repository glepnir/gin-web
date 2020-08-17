// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package global

import (
	"os"

	"github.com/glepnir/gin-web/pkg/auth"
	"github.com/glepnir/gin-web/pkg/auth/jwtauth"
)

func NewAuth() auth.Auther {
	pwd, _ := os.Getwd()
	withexpired := jwtauth.WithExpired(7200)
	withprivate := jwtauth.WithPrivateKey(pwd + "/configs/app.rsa")
	withpublic := jwtauth.WithPubKeyPath(pwd + "/configs/app.rsa.pub")
	auth := jwtauth.NewJwtAuth(withexpired, withprivate, withpublic)
	return auth
}
