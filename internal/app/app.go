// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glepnir/gin-web/internal/config"
	"github.com/glepnir/gin-web/internal/middlewares"
	"github.com/glepnir/gin-web/internal/routes"
	"github.com/glepnir/gin-web/internal/storage"
)

type Application struct {
	Config config.Config
	Route  *gin.Engine
}

func NewApplication(route *gin.Engine) *Application {
	return &Application{Route: route}
}

func (a *Application) CreateApp() {
	a.Config.MustLoadConf()
	a.Route.Use(gin.Recovery())
	configureDataBase(a.Config.DataBase)
	configureRouter(a.Route)
}

func (a *Application) Run() {
	addr := a.Config.HTTP.Host + ":" + a.Config.HTTP.Port
	shutdowntimeout := a.Config.HTTP.ShutDownTimeOut
	srv := &http.Server{
		Addr:    addr,
		Handler: a.Route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	defer storage.CloseDB()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Fatal("server shut down")
	ctx, cancel := context.WithTimeout(context.Background(), shutdowntimeout*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown err:", err)
	}

	log.Fatal("server exit")
}

func configureDataBase(dbconfig config.DataBase) {
	err := storage.NewDB(dbconfig)
	if err != nil {
		panic(err)
	}
}

func configureRouter(r *gin.Engine) {
	r.LoadHTMLGlob("template/*")
	r.Static("/static/", "./static/")
	g := r.Group("/")
	routes.HealthRouter(g)
	routes.LoginRoute(g)
	g.Use(middlewares.CheckAuth())
	g.Use(middlewares.CheckPermission())
	g.GET("renderusers", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-list.html", nil)
	})
	g.GET("rendercreateuser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-add.html", nil)
	})
	routes.UserRoute(g)
	routes.RoleRoute(g)
	routes.IndexRoute(g)
}
