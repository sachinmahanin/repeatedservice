package web

import "github.com/sachinmahanin/passwordRepeated/handler/miscellaneous"

// func pointers for injection / testing: web.go
var (
	miscellaneousSwaggerHandler  = miscellaneous.SwaggerHandler
	miscellaneousSwaggerRedirect = miscellaneous.SwaggerRedirect
	registeredUtilityRoutesFunc  = registeredUtilityRoutes
	registeredBusinessRoutesFunc = registeredBusinessRoutes
)
