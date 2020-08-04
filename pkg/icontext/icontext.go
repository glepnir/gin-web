// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package icontext

import (
	"context"

	"github.com/gin-gonic/gin"
)

type contextInject struct {
	context.Context
	ctx *gin.Context
}

func NewContextInject(ctx *gin.Context) context.Context {
	return &contextInject{context.Background(), ctx}
}
