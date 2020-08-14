// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"time"

	"github.com/glepnir/gin-web/internal/config"
	"github.com/glepnir/gin-web/internal/storage/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

type DB struct{}

func NewDB(storage config.DataBase) error {
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
		return err
	}

	if err := conn.DB().Ping(); err != nil {
		return err
	}

	conn.DB().SetMaxIdleConns(storage.MaxIdle)
	conn.DB().SetMaxOpenConns(storage.MaxConn)
	conn.DB().SetConnMaxLifetime(time.Duration(storage.MaxLifeTime) * time.Minute)
	db := &DB{}
	db.Set(conn)
	AutoMigrate(conn, storage.TablePrefix)
	return nil
}

func (d *DB) Get() *gorm.DB {
	return conn
}

func (d *DB) Set(db *gorm.DB) {
	conn = db
}

func CloseDB() {
	defer conn.Close()
}

func AutoMigrate(conn *gorm.DB, tableprefix string) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tableprefix + defaultTableName
	}

	if !conn.HasTable("users") {
		conn.AutoMigrate(&entity.User{})
	}
	if !conn.HasTable("menus") {
		conn.AutoMigrate(&entity.Menu{})
	}
	if !conn.HasTable("roles") {
		conn.AutoMigrate(&entity.Role{})
	}
}
