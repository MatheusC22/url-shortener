package main

import (
	"goAPI/configs"
	"goAPI/routes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func router01() http.Handler {
	//gin.SetMode(gin.ReleaseMode)
	r1 := gin.New()
	r1.Use(gin.Recovery())
	routes.UrlRoutes(r1)
	routes.UserRoutes(r1)
	routes.AuthRoutes(r1)
	r1.GET("/ping", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Pong",
				"Server":  01,
				"URL":     c.FullPath(),
			},
		)
	})

	return r1
}

func router02() http.Handler {
	//gin.SetMode(gin.ReleaseMode)
	r2 := gin.New()
	r2.Use(gin.Recovery())
	routes.UrlRoutes(r2)
	routes.UserRoutes(r2)
	routes.AuthRoutes(r2)
	r2.GET("/ping", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Pong",
				"Server":  02,
				"URL":     c.Request.URL.Path,
			},
		)
	})

	return r2
}

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	apiConfig := configs.GetApi()

	server01 := &http.Server{
		Addr:         apiConfig.Port_1,
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         apiConfig.Port_2,
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
