// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/app"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func initTestCreateUser(body map[string]interface{}) (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := app.NewApplication(r)
	app.CreateApp()

	w := httptest.NewRecorder()
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/create", strings.NewReader(string(b)))
	req.Header.Set("Content-type", "application/json")

	r.ServeHTTP(w, req)
	return w, r

}

func TestCreateUser(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		body := map[string]interface{}{
			"username": "tester",
			"email":    "test@test.org",
			"password": "123456",
		}
		actual := Response{}
		w, _ := initTestCreateUser(body)
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Create user success", actual.Message)
		assert.Equal(t, http.StatusOK, actual.Status)
		assert.Empty(t, actual.Data)
	})
}
