// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

type TokenInfo struct {
	AccessToken  string
	RefreshToken string
}

type Auther interface {
	LoadRsaFile() error
	GenerateToken(userID string) (TokenInfo, error)
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken()
}
