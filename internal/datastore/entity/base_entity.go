// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CreateAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}

// BeforeCreate is hooks to create uuid
func (b *Base) BeforCreate(scope *gorm.Scope) error {
	uuid := uuid.New().String()
	return scope.SetColumn("ID", uuid)
}
