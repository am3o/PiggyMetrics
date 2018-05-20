package main

import (
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// ConfigurationListenKey discribes the key-word for the configuration file.
	ConfigurationListenKey = "service.listen"
	// ServiceName discribes the name of the service.
	ServiceName = "Account-Service"
)

var (
	// ConfigurationsName discribes the state of the service
	ConfigurationsName = "debug"
	// ConfigurationFilePath discribes the location of the configuration file.
	ConfigurationFilePath = "configuration/account"
	// VersionNumber discribe the version of the service
	VersionNumber = "0.0.0"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithFields(logrus.Fields{
		"configuration": ConfigurationsName,
		"version":       VersionNumber,
		"pid":           os.Getpid(),
	}).Infof("Starting account server")
}

func main() {
	server, err := initializeServer()
	if err != nil {
		logrus.Error(err)
		return
	}

	server.Run(viper.GetString(ConfigurationListenKey))
}

// initializeServer creates all pertinent endpoints of the service and returns the server.
func initializeServer() (engine *gin.Engine, err error) {
	viper.SetConfigName("configuration")
	viper.AddConfigPath(ConfigurationFilePath)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(err)
		return
	}

	engine = gin.Default()

	internal := engine.Group("/internal")
	{
		internal.GET("/health", HealthCheck)
		internal.GET("/version", VersionCheck)
	}

	return engine, nil
}

// VersionCheck returns the current version of the sevice.
func VersionCheck(c *gin.Context) {
	c.String(http.StatusOK, VersionNumber)
}

// HealthCheck returns the current health state of the service.
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
