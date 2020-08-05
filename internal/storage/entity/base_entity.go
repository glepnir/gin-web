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
	CreateAt time.Time `gorm:"type:datatime" json:"create_at"`
	UpdateAt time.Time `gorm:"type:datatime" json:"update_at"`
}

// BeforeCreate is hooks to create uuid
func (b *Base) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}
