package cusomization

import (
	"time"

	logger "github.com/sachinmahanin/passwordrepeated/middleware/logger"
	"github.com/sachinmahanin/passwordrepeated/sessionutil"
	web "github.com/sachinmahanin/passwordrepeated/web"
	webserver "github.com/zhongjie-cai/web-server"
)

// func pointers for injection / testing: customization.go
var (
	webRegisteredRoutes  = web.RegisteredRoutes
	webRegisteredStatics = web.RegisteredStatics
	timeParseDuration    = time.ParseDuration
	customizeLoggingFunc = logger.CustomizeLoggingFunc
	webserverWrapError   = webserver.WrapError
	prepareSessionFunc   = sessionutil.PrepareSession
)
