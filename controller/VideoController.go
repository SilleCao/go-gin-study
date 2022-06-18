package controller

import (
	"net/http"

	"github.com/EDDYCJY/go-gin-example/entity"
	"github.com/EDDYCJY/go-gin-example/service"
	"github.com/EDDYCJY/go-gin-example/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IVideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type VideoController struct {
	service service.IVideoService
}

var validate *validator.Validate

func New(service service.IVideoService) IVideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &VideoController{
		service: service,
	}
}

func (vc *VideoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}

func (vc *VideoController) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	// ctx.BindJSON(&video)
	vc.service.Save(video)
	return nil
}

func (vc *VideoController) ShowAll(ctx *gin.Context) {
	videos := vc.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
