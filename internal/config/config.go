// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/glepnir/gin-web/pkg/finder"
	"github.com/spf13/viper"
)

var (
	once sync.Once
)

type Config struct {
	HTTP     HTTP
	DataBase DataBase
}

type HTTP struct {
	Host string
	Port string
}

type DataBase struct {
	Driver      string
	User        string
	Password    string
	Host        string
	Port        int
	DBname      string
	Charset     string
	MaxIdle     int
	MaxConn     int
	MaxLifeTime int
}

func (c *Config) MustLoadConf() {
	once.Do(func() {
		root, err := finder.InferRootDir("configs")
		if err != nil {
			panic(err)
		}
		ConfPath := root + string(os.PathSeparator) + "configs"

		viper.SetConfigName("config")
		viper.AddConfigPath(ConfPath)

		err = viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		viper.WatchConfig()
	})

	c.HTTP = HTTP{
		Host: viper.GetString("http.host"),
		Port: viper.GetString("http.port"),
	}

	c.DataBase = DataBase{
		Driver:      viper.GetString("storage.driver"),
		User:        viper.GetString("storage.user"),
		Password:    viper.GetString("storage.password"),
		Host:        viper.GetString("storage.host"),
		Port:        viper.GetInt("storage.port"),
		DBname:      viper.GetString("storage.dbname"),
		Charset:     viper.GetString("storage.charset"),
		MaxIdle:     viper.GetInt("storage.max_idle"),
		MaxConn:     viper.GetInt("storage.max_conn"),
		MaxLifeTime: viper.GetInt("storage.max_lifetime"),
	}

}
