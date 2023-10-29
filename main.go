package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"osa.synapsis.chalange/utils/db"
	"osa.synapsis.chalange/utils/routes"
)

func main() {
	//setup environment
	viper.SetConfigName("app")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	//connect db
	db, err := db.NewConnectDB()
	if err != nil {
		log.Fatalln(err.Error())
	}

	//setup fiber
	app := fiber.New(fiber.Config{
		AppName: "Online Shop Application",
		ServerHeader: "Go Fiber",
	})

	routes.NewRoutes(app, db)
	log.Fatal(app.Listen(
		fmt.Sprintf(":%d", viper.GetInt("PORT_SERVER"))))
}