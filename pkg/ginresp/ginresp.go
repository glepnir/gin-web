// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ginresp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnAuthorized(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    "-1",
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func BadRequest(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    "-1",
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func Conflict(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusConflict, gin.H{
		"code":    "-1",
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func NotFound(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    "-1",
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func Ok(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func OkWithCount(c *gin.Context, message string, data interface{}, count int, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": message,
		"count":   count,
		"data":    data,
		"error":   err,
	})
}
func PassWordWrong(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "-1",
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func Forbidden(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusForbidden, gin.H{
		"code":    "-1",
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func InternalError(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    "-1",
		"message": message,
		"data":    data,
		"error":   err,
	})
}
