package main

import (
	"github.com/agustinrabini/api-items-project/cmd/api/app"
	"github.com/agustinrabini/api-items-project/cmd/api/config"

	_ "github.com/agustinrabini/api-items-project/docs"
)

// @title Jopit API Items
// @version 1.0
// @description This is a jopit api items.
// @termsOfService http://swagger.io/terms/

// @contact.name Matias Nuñez
// @contact.url http://www.swagger.io/support
// @contact.email matiasne45@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	config.Load()
	app.Start()
}
