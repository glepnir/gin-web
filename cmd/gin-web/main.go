// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/app"
)

func main() {
	r := gin.Default()
	instance := app.NewApplication(r)
	instance.CreateApp()
	instance.Run()
}
