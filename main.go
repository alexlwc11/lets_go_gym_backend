package main

import (
	"lets_go_gym_backend/configs"
	"log"
)

//	@title						Let's go gym API
//	@version					1.0
//	@description				Let's go gym API endpoints.
//
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//
//	@host						localhost:8080
//	@BasePath					/api/v1
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type in format of "Bearer --TOKEN--".

func main() {
	// load the database config
	mConfig := configs.LoadConfig()

	// init database connection
	sqlStorage := NewSQLStorage(mConfig)
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	// init router
	router := NewAPIServer("0.0.0.0:8080", db)
	router.Init()
	router.SetupSwagger()
	router.Run()
}
