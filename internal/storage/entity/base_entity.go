// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

type Base struct {
	ID        xid.ID     `gorm:"column:id;primary_key;"`
	CreatedAt time.Time  `gorm:"column:created_at;index;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;index;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;"`
}

// BeforeCreate is hooks to create uuid
func (b *Base) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", xid.New())
}
