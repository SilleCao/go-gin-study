package main

import (
	"io"
	"net/http"
	"os"

	"github.com/EDDYCJY/go-gin-example/controller"
	"github.com/EDDYCJY/go-gin-example/middlewares"
	"github.com/EDDYCJY/go-gin-example/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoServcie    service.IVideoService       = service.New()
	videoController controller.IVideoController = controller.New(videoServcie)

	loginService    service.ILoginService      = service.NewLoginService()
	jwtService      service.IJWTService        = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	// server := gin.Default()

	server := gin.New()

	server.Static("/css", "./template/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "SUCCESS",
			})
		})

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			// ctx.JSON(200, videoController.Save(ctx))
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video Input is Valid!!",
				})
			}
		})
	}

	viewRoutes := server.Group("/view", middlewares.BasicAuth())
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
