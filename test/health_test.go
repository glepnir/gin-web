// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestCreateApp(t *testing.T) {
	r := gin.Default()
	app := app.NewApplication(r)
	app.CreateApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/health", nil)
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
