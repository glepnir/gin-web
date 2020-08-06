// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ginresp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusOK,
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func Conflict(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusConflict, gin.H{
		"status":  http.StatusConflict,
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func NotFound(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  http.StatusNotFound,
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func Ok(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func Forbidden(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusForbidden, gin.H{
		"status":  http.StatusForbidden,
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func InternalError(c *gin.Context, message string, data interface{}, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"message": message,
		"data":    data,
		"error":   err,
	})
}
