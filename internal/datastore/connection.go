// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datastore

import (
	"fmt"
	"time"

	"github.com/glepnir/gin-web/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB struct {
	conn *gorm.DB
}

func NewDB(storage config.Storage) (*gorm.DB, error) {
	driver := storage.Driver

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		storage.User,
		storage.Password,
		storage.Host,
		storage.Port,
		storage.DBname,
		storage.Charset)

	conn, err := gorm.Open(driver, dbURI)
	if err != nil {
		return nil, err
	}

	if err := conn.DB().Ping(); err != nil {
		return nil, err
	}

	db := DB{conn}
	conn.DB().SetMaxIdleConns(storage.MaxIdle)
	conn.DB().SetMaxOpenConns(storage.MaxConn)
	conn.DB().SetConnMaxLifetime(time.Duration(storage.MaxLifeTime) * time.Minute)
	return db.conn, nil
}
