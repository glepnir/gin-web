// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jwtauth

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/glepnir/gin-web/pkg/auth"
)

type JwtAuth struct {
	verifyKey               *rsa.PublicKey
	signKey                 *rsa.PrivateKey
	privKeyPath, pubKeyPath string
	expired                 time.Duration
}

var _ auth.Auther = (*JwtAuth)(nil)

type JwtAuthFunc func(*JwtAuth)

func WithPrivateKey(privKeyPath string) JwtAuthFunc {
	return func(j *JwtAuth) {
		j.privKeyPath = privKeyPath
	}
}

func WithPubKeyPath(pubKeyPath string) JwtAuthFunc {
	return func(j *JwtAuth) {
		j.pubKeyPath = pubKeyPath
	}
}

func WithExpired(expired int) JwtAuthFunc {
	return func(j *JwtAuth) {
		j.expired = time.Duration(expired)
	}
}

func NewJwtAuth(opts ...JwtAuthFunc) auth.Auther {
	j := new(JwtAuth)
	for _, o := range opts {
		o(j)
	}
	return j
}

func (j *JwtAuth) LoadRsaFile() error {
	verifyBytes, err := ioutil.ReadFile(j.pubKeyPath)

	if err != nil {
		return err
	}

	j.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}

	signBytes, err := ioutil.ReadFile(j.privKeyPath)
	if err != nil {
		return err
	}

	j.signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}
	return nil
}

func (j *JwtAuth) GenerateToken(userID string) (auth.TokenInfo, error) {
	j.LoadRsaFile()
	t := auth.TokenInfo{}
	accesstoken, err := j.GenerateAccessToken(userID)
	t.AccessToken = accesstoken
	if err != nil {
		return t, err
	} else {
		return t, nil
	}
}

func (j *JwtAuth) GenerateAccessToken(userID string) (string, error) {
	j.LoadRsaFile()
	now := time.Now()
	expiresAt := now.Add(j.expired * time.Second).Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   userID,
	})

	tokenstring, err := t.SignedString(j.signKey)

	if err != nil {
		return tokenstring, err
	}
	return tokenstring, nil
}

func (j *JwtAuth) GenerateRefreshToken() {

}

func (j *JwtAuth) parseToken(tokenstring string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.verifyKey, nil
	})
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims.(*jwt.StandardClaims), nil
}

func (j *JwtAuth) ParseUserID(tokenstring string) (string, error) {
	j.LoadRsaFile()
	claims, err := j.parseToken(tokenstring)
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}
