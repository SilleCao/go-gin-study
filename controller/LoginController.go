package controller

import (
	"github.com/EDDYCJY/go-gin-example/dto"
	"github.com/EDDYCJY/go-gin-example/service"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService service.ILoginService
	jwtService   service.IJWTService
}

func NewLoginController(loginService service.ILoginService, jwtService service.IJWTService) LoginController {
	return LoginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (crtl LoginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := crtl.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return crtl.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
