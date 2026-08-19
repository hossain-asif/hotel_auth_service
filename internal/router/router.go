package router

import (
	"github.com/hossain-asif/hotel_auth_service/internal/pkg/module"
	"github.com/hossain-asif/hotel_auth_service/internal/user"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

var Modules = []module.Module{
	&user.UserModule{},
	// &role.RoleModule{},
}
