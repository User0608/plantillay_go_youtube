package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/user0608/api_rest/api/router"
	"github.com/user0608/api_rest/cmd/injectors"
	"github.com/user0608/api_rest/configs"
)

func main() {
	serverConfigs, err := configs.LoadServiceConfigs("service_config.json")
	if err != nil {
		log.Fatalln("No se cargo las configuraciones de servicio:", err.Error())
	}
	log.Println("Server configs OK!")

	if err := injectors.LoadConfig("db_config.json"); err != nil {
		log.Fatalln("Error base de datos,", err.Error())
	}
	log.Println("Base de datos OK!")
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: serverConfigs.Cors.AllowOrigins,
		AllowMethods: serverConfigs.Cors.AllowMethods,
	}))
	server.HideBanner = true
	router.Updrade(server)
	if err := server.Start(serverConfigs.Address); err != nil {
		log.Fatal(err)
	}
	log.Println("Ok!")
}
