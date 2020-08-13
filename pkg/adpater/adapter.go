// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adpater

import (
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

var _ persist.Adapter = (*CasbinAdapter)(nil)

type CasbinAdapter struct {
}

func (c *CasbinAdapter) LoadPolicy(m casbinmodel.Model) error {
	return nil

}

func (c *CasbinAdapter) SavePolicy(m casbinmodel.Model) error {
	return nil
}

func (a *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

func (a *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

func (a *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
