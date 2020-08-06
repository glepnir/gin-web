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
	timeout := a.Config.HTTP.TimeOut
	srv := &http.Server{
		Addr:    addr,
		Handler: a.Route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Fatal("server shut down")
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown err:", err)
	}

	log.Fatal("erver exit")
}

func configureDataBase(cdb config.DataBase) {
	err := storage.NewDB(cdb)
	if err != nil {
		panic(err)
	}
}

func configureRouter(r *gin.Engine) {
	g := r.Group("/api/v1")
	routes.HealthRouter(g)
	routes.UserRoute(g)
}
