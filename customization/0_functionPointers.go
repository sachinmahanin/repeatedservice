package cusomization

import (
	"time"

	logger "github.com/sachinmahanin/passwordRepeated/middleware/logger"
	"github.com/sachinmahanin/passwordRepeated/sessionutil"
	web "github.com/sachinmahanin/passwordRepeated/web"
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
