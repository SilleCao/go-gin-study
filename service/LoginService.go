package service

type ILoginService interface {
	Login(username string, password string) bool
}

type LoginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() ILoginService {
	return &LoginService{
		authorizedUsername: "Sille",
		authorizedPassword: "123456",
	}
}

func (service *LoginService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
