package main

import (
	"isjhar/template/echo-golang/utils"
	"isjhar/template/echo-golang/view/routers"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}
	e.Use(middleware.Recover())

	environment := utils.GetEnvironment()
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	if environment != utils.DEVELOPMENT {
		output := &utils.CustomLogger{}
		log.SetOutput(output)
		utils.LogLevel = utils.LogWarningLevel
	}

	//CORS Config
	CORSConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}
	e.Use(middleware.CORSWithConfig(CORSConfig))

	routers.Route(e)

	e.Logger.Fatal(e.Start(":1323"))
}
