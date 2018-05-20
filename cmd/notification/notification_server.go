package main

import (
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// ConfigurationListenKey describes the key-word for the configuration file.
	ConfigurationListenKey = "service.listen"
	// ServiceName describes the name of the service.
	ServiceName = "Notification-Service"
)

var (
	// ConfigurationsName describes the state of the service
	ConfigurationsName = "debug"
	// ConfigurationFilePath describes the location of the configuration file.
	ConfigurationFilePath = "configuration/noftification"
	// VersionNumber describes the version of the service
	VersionNumber = "0.0.0"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithFields(logrus.Fields{
		"configuration": ConfigurationsName,
		"version":       VersionNumber,
		"pid":           os.Getpid(),
	}).Infof("Starting notification server")

	viper.SetConfigName("configuration")
	viper.AddConfigPath(ConfigurationFilePath)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Warnf("couldn't load configuration: %v", err)
	}
}

func main() {
	server, err := initializeNotificationServer()
	if err != nil {
		logrus.Error(err)
		return
	}

	server.Run(viper.GetString(ConfigurationListenKey))
}

// initializeNotificationServer creates all pertinent endpoints of the service and returns the server.
func initializeNotificationServer() (engine *gin.Engine, err error) {
	engine = gin.Default()

	internal := engine.Group("/internal")
	{
		internal.GET("/health", HealthCheck)
		internal.GET("/version", VersionCheck)
	}

	return engine, nil
}

// VersionCheck returns the current version of the service.
func VersionCheck(c *gin.Context) {
	c.String(http.StatusOK, VersionNumber)
}

// HealthCheck returns the current health state of the service.
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
