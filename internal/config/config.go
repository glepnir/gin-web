// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
)

type Config struct {
	HTTP    HTTP
	Storage Storage
}

type HTTP struct {
	Host string
	Port string
}

type Storage struct {
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
		ConfPath := inferRootDir() + string(os.PathSeparator) + "configs"

		viper.SetConfigName("config")
		viper.AddConfigPath(ConfPath)

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		viper.WatchConfig()
	})

	c.HTTP = HTTP{
		Host: viper.GetString("http.host"),
		Port: viper.GetString("http.port"),
	}

	c.Storage = Storage{
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

func inferRootDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(d string) string
	infer = func(d string) string {
		if d == "/" {
			panic("Please run in the root of project" + cwd)
		}
		_, err := os.Stat(d + string(os.PathSeparator) + "configs")
		if err == nil || os.IsExist(err) {
			return d
		}
		return infer(filepath.Dir(d))
	}
	return infer(cwd)
}
