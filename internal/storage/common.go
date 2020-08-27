// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import "github.com/jinzhu/gorm"

func FindOne(conn *gorm.DB, model, out, arg interface{}, query ...string) (interface{}, bool) {
	if len(query) > 2 {
		exist := conn.Model(model).Select(query[0]).Where(query[1], arg).First(out).RecordNotFound()
		if exist {
			return out, false
		} else {
			return out, true
		}
	}
	return nil, false
}

func CreateOne(conn *gorm.DB, model interface{}) error {
	tx := conn.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(model).Error; err != nil {
		tx.Rollback()
		return nil
	}
	return tx.Commit().Error
}

func Exist(conn *gorm.DB, args ...interface{}) bool {
	if len(args) == 3 {
		exist := conn.Select(args[0]).Where(args[1], args[2]).RecordNotFound()
		if exist {
			return false
		}
		return true
	}
	return false
}
