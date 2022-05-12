// Package factories provides helper functions for initializing services. Because of complex inter-service dependencies, it is easier to centralize this all in one place.
package factories

import (
	"github.com/seanvm/serverless-golang-starter-backend/app"
	"github.com/seanvm/serverless-golang-starter-backend/app/database"
)

type ServiceFactory struct {
	Db *database.Datastore
}

func (s *ServiceFactory) BuildUserService() app.UserServicer {
	return app.NewUserService(s.Db.UserRepository)
	// return app.NewUserService(s.Db.UserRepository)
}
