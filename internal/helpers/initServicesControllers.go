package helpers

import (
	"api-integracao/internal/controllers"
	"api-integracao/internal/service"

	"github.com/couchbase/gocb/v2"
)

func InitControllers(scope *gocb.Scope) controllers.ControllerInitializaer {
	services := service.InitServices(scope)
	controllers := controllers.InitControllers(*services)
	return controllers
}
