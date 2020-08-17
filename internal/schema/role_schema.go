// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schema

import (
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/glepnir/gin-web/internal/storage"
)

type RoleSchema struct {
	RoleName string `json:"rolename" validate:"required"`
	Path     string `json:"path" validate:"required"`
	Method   string `json:"method" validate:"required"`
}

func LoadCasbin() *casbin.Enforcer {
	dbURI := storage.GetDBURI()
	a, _ := gormadapter.NewAdapter("mysql", dbURI, true) // Your driver and data source.
	rootpath, _ := os.Getwd()
	path := rootpath + "/configs/model.conf"
	e, _ := casbin.NewEnforcer(path, a)
	e.LoadPolicy()
	return e
}

func (r *RoleSchema) AddRole() (bool, error) {
	e := LoadCasbin()
	result, err := e.AddPolicy(r.RoleName, r.Path, r.Method)
	return result, err
}

func (r *RoleSchema) DeleteRole() {
	e := LoadCasbin()
	e.RemovePolicy(r.RoleName, r.Path, r.Method)
}
