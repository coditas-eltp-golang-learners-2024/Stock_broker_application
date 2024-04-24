#!/bin/bash

if [ $# -ne 3 ]; then
    echo "Usage: $0 <servicename> <port> <goversion>"
    exit 1
fi

microservice=$1
port=$2
goversion=$3

# Modify go.work
sed -i "/use (/a \ \ \ \ .\/src\/app\/${microservice}" go.work
echo "microservice entry added successfully in go.work file"

# Create directory structure inside src/app/
mkdir -p src/app/${microservice}/{handler,business,repositories,tests,router,models,commons/constants}
echo "microservice folders created"

# Tests
cat <<EOT > src/app/${microservice}/tests/start_test.go
package tests

import (
    "context"
    setupTest "stock_broker_application/src/setupTest"
    "testing"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/suite"
)

type ${microservice} struct {
    suite.Suite
    ctrl *gomock.Controller
    ctx context.Context
}

func (suite *${microservice}) SetupSuite() {
    setupTest.InitSuiteConfigs()
}

func (suite *${microservice}) BeforeTest(suiteName, testName string) {
    suite.ctx = context.Background()
    suite.ctrl = gomock.NewController(suite.T())
}

func (suite *${microservice}) AfterTest(suiteName, testName string) {
    suite.ctrl.Finish()
}

func Test${microservice}(t *testing.T) {
    suite.Run(t, new(${microservice}))
}
EOT
echo "microservice test cases file created"

# Router
cat <<EOT > src/app/${microservice}/router/routes.go
package router

import (
    "net/http"
    "time"

    serviceConstant "${microservice}/commons/constants"
    genericConstants "stock_broker_application/src/constants"
    dbDeviceRepository "stock_broker_application/src/database/repository"
    "stock_broker_application/src/middleware/authorization"
    "stock_broker_application/src/middleware/cryptoHandshake"
    rateLimitMiddleware "stock_broker_application/src/middleware/ginRateLimiter"
    headerMiddleware "stock_broker_application/src/middleware/headerCheck"
    metricsMiddleware "stock_broker_application/src/middleware/metrics"
    jwtUtils "stock_broker_application/src/utils/authorization"
    "stock_broker_application/src/utils/configs"
    "stock_broker_application/src/utils/metrics"

    prometheusHandler "github.com/prometheus/client_golang/prometheus/promhttp"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/juju/ratelimit"
)

func init() {
    gin.SetMode(gin.ReleaseMode)
}

// GetRouter is used to get the router configured with the middlewares and the routes
func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
    router := gin.New()
    router.Use(middlewares...)
    router.Use(gin.Recovery())

    applicationConfig := configs.GetApplicationConfig()
    limiter := ratelimit.NewBucket(time.Duration(applicationConfig.AppConfig.RateLimitIntervalInSecond)*time.Second,
        int64(applicationConfig.AppConfig.RateLimitRequestPerInterval))

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{genericConstants.PostMethod, genericConstants.GetMethod, genericConstants.PutMethod, genericConstants.DeleteMethod, genericConstants.PatchMethod},
        AllowHeaders:     []string{genericConstants.AllowHeaderOriginConfig, genericConstants.Authorization},
        ExposeHeaders:    []string{genericConstants.ExposeHeaderContentLengthConfig},
        AllowCredentials: true,
        MaxAge:           300 * time.Second,
    }))

    metrics.Init()
    router.GET(serviceConstant.Metrics, gin.WrapH(prometheusHandler.Handler()))

    enableUIBFFEncDec := applicationConfig.AppConfig.EnableUIBFFEncDec
    enableRateLimit := applicationConfig.AppConfig.EnableRateLimit

    jwtUtils := jwtUtils.NewJwtTokenUtils()
    jwtMiddleware := authorization.AuthorizeJWtToken(jwtUtils)

    dbDeviceRepository := dbDeviceRepository.NewDeviceRepository()
    encryptMiddleware := cryptoHandshake.EncryptMiddleware(dbDeviceRepository)
    decryptMiddleware := cryptoHandshake.DecryptionMiddleware(dbDeviceRepository)

    v1Routes := router.Group(genericConstants.RouterV1Config)
    {
        v1Routes.GET(serviceConstant.${microservice^}HealthCheck, func(c *gin.Context) {
            response := map[string]string{
                genericConstants.ResponseMessageKey: genericConstants.BFFResponseSuccessMessage,
            }
            c.JSON(http.StatusOK, response)
        })
        if enableRateLimit {
            v1Routes.Use(rateLimitMiddleware.RateLimitMiddleware(limiter))
        }
        v1Routes.Use(headerMiddleware.HeaderCheck(serviceConstant.ServiceName), metricsMiddleware.Metric(), jwtMiddleware)
        if enableUIBFFEncDec {
            v1Routes.Use(encryptMiddleware, decryptMiddleware)
        }
    }
    return router
}
EOT
echo "microservice router file created"

# Constants
cat <<EOT > src/app/${microservice}/commons/constants/constants.go
package constants

// ${microservice} NEST API URL Keys
const (
    ServiceName        = "${microservice}"
    PortDefaultValue   = ${port}
)
EOT
echo "microservice constants file created"


# Constants - Errors
cat <<EOT > src/app/${microservice}/commons/constants/errors.go
package constants

// Add your error constants here
EOT
echo "microservice error constants file created"

# Constants - Routes
cat <<EOT > src/app/${microservice}/commons/constants/routes.go
package constants

// ${microservice} route Constants
const (
    SwaggerRoute                = "/swagger/*any"
    ${microservice^}HealthCheck = "/${microservice}/health-check"
    Metrics                     = "/metrics"
)
EOT
echo "microservice routes constants file created"

# Commons
cat <<EOT > src/app/${microservice}/commons/common.go
package commons

// Add your common functionalities here
EOT
echo "microservice common utils function file created"

# Main
cat <<EOT > src/app/${microservice}/main.go
package main

import (
    "context"
    "fmt"
    genericConstants "stock_broker_application/src/constants"
    loggerMiddleware "stock_broker_application/src/middleware/logger"
    "stock_broker_application/src/utils/configs"
    "stock_broker_application/src/utils/flags"
    "stock_broker_application/src/utils/logger"
    "stock_broker_application/src/utils/postgres"
    serviceConstants "${microservice}/commons/constants"
    "${microservice}/router"

    "go.uber.org/zap"
)

func main() {
    ctx := context.Background()
    initConfigs(ctx)
    configs.SetHostName(serviceConstants.ServiceName)
    initLogger(ctx)
    log := logger.GetLoggerWithoutContext()
    err := postgres.InitPostgresDBConfig(ctx)
    if err != nil {
        log.With(zap.Error(err)).Fatal(genericConstants.PostgresDBInitializationError)
    }
    defer postgres.ClosePostgres(ctx)
    err = configs.InitApplicationConfigs(ctx)
    if err != nil {
        log.With(zap.Error(err)).Fatal(genericConstants.ConfigBindingFailedError)
        panic(fmt.Errorf(genericConstants.ConfigBindingFailedError))
    }
    configs.InitRegexPatterns()
    utils.InitApiUrls()
    startRouter(ctx)
}

func initConfigs(ctx context.Context) {
    configs.Init([]string{flags.BaseConfigPath(), flags.MockConfigPath()})
}

func initLogger(ctx context.Context) {
    LoggerConfig, err := configs.Get(genericConstants.LoggerConfig)
    if err != nil {
        panic(err)
    }
    logger.SetupLogging(LoggerConfig.GetString(genericConstants.LogLevelKey))
}

func startRouter(ctx context.Context) {
    log := logger.GetLoggerWithoutContext()
    router := router.GetRouter(loggerMiddleware.Logger())
    log.Info(fmt.Sprintf(genericConstants.RunningServerPort, serviceConstants.PortDefaultValue))
    err := router.Run(fmt.Sprintf(":%d", serviceConstants.PortDefaultValue))
    if err != nil {
        log.With(zap.Error(err)).Fatal(genericConstants.ExternalServiceError)
    }
}
EOT
echo "microservice main file created"

# Run go mod init command
cd src/app/${microservice}
go mod init ${microservice}
go mod edit -go=${goversion}
go mod tidy

echo "service structure for ${microservice} created successfully"
