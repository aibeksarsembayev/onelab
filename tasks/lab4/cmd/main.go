// создать template crud используя фреймворк echo, (4 роута, разные методы http, разные способы передачи данных - by query, by body, etc.)
// завернуть в докер контейнер, и запустить локально
// проверить роуты апи через постман либо через запросы с другого го инстанса (net/http)
// загрузить лабку на гитхаб/гитлаб
package main

import (
	"time"

	_userHttpDelivery "github.com/aibeksarsembayev/onelab/tasks/lab4/user/handlers/http"
	_userRepo "github.com/aibeksarsembayev/onelab/tasks/lab4/user/repository"
	_userUsecase "github.com/aibeksarsembayev/onelab/tasks/lab4/user/usecases"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	userRepo := _userRepo.NewDBUserRepository() // pass db conn

	timeoutContext := time.Duration(5 * time.Second)

	uUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	_userHttpDelivery.NewUserHandler(e, uUsecase)

	e.Logger.Fatal(e.Start(":8080"))

}
