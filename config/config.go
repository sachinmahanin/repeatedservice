package config

import (
	"container/list"

	webserver "github.com/zhongjie-cai/web-server"
)

var (
	// AppVersion returns the version information of the application
	AppVersion = "1.0"

	// AppPort returns the hosting port of the application
	AppPort = "18605"

	// AppName returns the name of the application
	AppName = "service-repeated-password"

	// AppPath returns the execution path of the application
	AppPath = "."

	// DefaultNetworkTimeout returns the default network timeout value of the application
	DefaultNetworkTimeout = "3m"

	// HostName is the name of the current host machine
	HostName = ""

	// AllowedLogType returns the default allowed log type of the application
	AllowedLogType = webserver.LogTypeAppRoot

	// AllowedLogLevel returns the default allowed log level of the application
	AllowedLogLevel = webserver.LogLevelDebug

	// stores older passwords
	PasswordHistory = list.New()
)

// SetupApplication initiates all application related root configs
func SetupApplication() error {
	HostName = osGetenv("HOSTNAME")
	AppVersion = osGetenv("APP_VERSION")
	AppPort = osGetenv("APP_PORT")
	AllowedLogLevel = webserverNewLogLevel(osGetenv("ALLOWED_LOG_LEVEL"))
	AllowedLogType = webserverNewLogType(osGetenv("ALLOWED_LOG_TYPE"))
	return nil
}
