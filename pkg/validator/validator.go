// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *CustomValidator) LayInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}

func (c *CustomValidator) Validate(schema interface{}) error {
	c.LayInit()
}
