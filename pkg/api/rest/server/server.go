/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package server is to handle REST API
package server

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"crypto/subtle"
	"fmt"
	"os"

	"net/http"

	// Black import (_) is for running a package's init() function without using its other contents.
	_ "github.com/cloud-barista/mc-terrarium/pkg/config"
	_ "github.com/cloud-barista/mc-terrarium/pkg/logger"
	"github.com/rs/zerolog/log"

	"github.com/cloud-barista/mc-terrarium/pkg/api/rest/custommiddleware"
	"github.com/cloud-barista/mc-terrarium/pkg/api/rest/handler"
	"github.com/cloud-barista/mc-terrarium/pkg/api/rest/route"
	"github.com/cloud-barista/mc-terrarium/pkg/tofu"

	"github.com/spf13/viper"

	// REST API (echo)
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// echo-swagger middleware
	_ "github.com/cloud-barista/mc-terrarium/pkg/api/rest/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/cloud-barista/mc-terrarium/pkg/readyz"
)

//var masterConfigInfos confighandler.MASTERCONFIGTYPE

const (
	infoColor    = "\033[1;34m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
	debugColor   = "\033[0;36m%s\033[0m"
)

const (
	website = " https://github.com/cloud-barista/mc-terrarium"
	banner  = `    

 ██████╗ ███████╗ █████╗ ██████╗ ██╗   ██╗
 ██╔══██╗██╔════╝██╔══██╗██╔══██╗╚██╗ ██╔╝
 ██████╔╝█████╗  ███████║██║  ██║ ╚████╔╝ 
 ██╔══██╗██╔══╝  ██╔══██║██║  ██║  ╚██╔╝  
 ██║  ██║███████╗██║  ██║██████╔╝   ██║   
 ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═════╝    ╚═╝   

 Computing Infrastructure Migration Technology
 ________________________________________________`
)

// RunServer func start Rest API server

// @title Multi-Cloud Terrarium  REST API
// @version latest
// @description Multi-Cloud Terrarium (mc-terrarium) aims to provide an environment to enrich multi-cloud infrastructure.

// @contact.name API Support
// @contact.url http://cloud-barista.github.io
// @contact.email contact-to-cloud-barista@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /terrarium

// @securityDefinitions.basic BasicAuth
func RunServer(port string) {

	log.Info().Msg("Setting Tofu command utility")
	if err := tofu.LoadRunningStatusMap(); err != nil {
		log.Warn().Msg(err.Error())
	}

	defer func() {
		if err := tofu.SaveRunningStatusMap(); err != nil {
			log.Error().Err(err).Msg("Failed to save running status map")
		}
	}()

	log.Info().Msg("Setting mc-terrarium REST API server")

	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger()) // default logger middleware in echo

	// Custom logger middleware with zerolog
	e.Use(custommiddleware.Zerologger())

	// Recover middleware recovers from panics anywhere in the chain, and handles the control to the centralized HTTP error handler.
	e.Use(middleware.Recover())

	// limit the application to 20 requests/sec using the default in-memory store
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	// Add a request ID to the context of each request
	e.Use(middleware.RequestID())

	e.HideBanner = true
	//e.colorer.Printf(banner, e.colorer.Red("v"+Version), e.colorer.Blue(website))

	allowedOrigins := viper.GetString("api.allow.origins")
	if allowedOrigins == "" {
		log.Fatal().Msg("allow_ORIGINS env variable for CORS is " + allowedOrigins +
			". Please provide a proper value and source setup.env again. EXITING...")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{allowedOrigins},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Conditions to prevent abnormal operation due to typos (e.g., ture, falss, etc.)
	enableAuth := viper.GetString("api.auth.enabled") == "true"

	apiUser := viper.GetString("api.username")
	apiPass := viper.GetString("api.password")

	if enableAuth {
		e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
			// Skip authentication for some routes that do not require authentication
			Skipper: func(c echo.Context) bool {
				if c.Path() == "/terrarium/readyz" ||
					c.Path() == "/terrarium/httpVersion" {
					return true
				}
				return false
			},
			Validator: func(username, password string, c echo.Context) (bool, error) {
				// Be careful to use constant time comparison to prevent timing attacks
				if subtle.ConstantTimeCompare([]byte(username), []byte(apiUser)) == 1 &&
					subtle.ConstantTimeCompare([]byte(password), []byte(apiPass)) == 1 {
					return true, nil
				}
				return false, nil
			},
		}))
	}

	fmt.Println("\n \n ")
	fmt.Print(banner)
	fmt.Println("\n ")
	fmt.Println("\n ")
	fmt.Printf(infoColor, website)
	fmt.Println("\n \n ")

	// Route for system management
	swaggerRedirect := func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/terrarium/api/index.html")
	}
	e.GET("/terrarium/api", swaggerRedirect)
	e.GET("/terrarium/api/", swaggerRedirect)
	e.GET("/terrarium/api/*", echoSwagger.WrapHandler)
	// e.GET("/terrarium/swagger/*", echoSwagger.WrapHandler)
	// e.GET("/terrarium/swaggerActive", rest_common.RestGetSwagger)

	e.GET("/terrarium/readyz", handler.Readyz)
	e.GET("/terrarium/httpVersion", handler.HTTPVersion)
	e.GET("/terrarium/tofuVersion", handler.TofuVersion)

	// A group for Multi-cloud Network APIs which has /terrarium as prefix
	groupMultiCloudNetwork := e.Group("/terrarium")
	// Resource Group APIs
	route.RegisterRoutesForTestEnv(groupMultiCloudNetwork)
	route.RegisterRoutesForRG(groupMultiCloudNetwork)
	route.RegisterRoutesForVPN(groupMultiCloudNetwork)

	// Sample API group (for developers to add new API)
	groupSample := groupMultiCloudNetwork.Group("/sample")
	route.RegisterSampleRoutes(groupSample)

	selfEndpoint := viper.GetString("self.endpoint")
	apidashboard := " http://" + selfEndpoint + "/terrarium/api"

	if enableAuth {
		fmt.Println(" Access to API dashboard" + " (username: " + apiUser + " / password: " + apiPass + ")")
	}
	fmt.Printf(noticeColor, apidashboard)
	fmt.Println("\n ")

	// A context for graceful shutdown (It is based on the signal package)selfEndpoint := os.Getenv("SELF_ENDPOINT")
	// NOTE -
	// Use os.Interrupt Ctrl+C or Ctrl+Break on Windows
	// Use syscall.KILL for Kill(can't be caught or ignored) (POSIX)
	// Use syscall.SIGTERM for Termination (ANSI)
	// Use syscall.SIGINT for Terminal interrupt (ANSI)
	// Use syscall.SIGQUIT for Terminal quit (POSIX)
	gracefulShutdownContext, stop := signal.NotifyContext(context.TODO(),
		os.Interrupt, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	// Wait graceful shutdown (and then main thread will be finished)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		// Block until a signal is triggered
		<-gracefulShutdownContext.Done()

		fmt.Println("\n[Stop] mc-terrarium REST API server")
		log.Info().Msg("stopping mc-terrarium REST API server")
		ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Panic(err)
		}
	}(&wg)

	log.Info().Msg("starting mc-terrarium REST API server")
	port = fmt.Sprintf(":%s", port)
	readyz.SetReady(true)
	if err := e.Start(port); err != nil && err != http.ErrServerClosed {
		e.Logger.Panic("shuttig down the server")
	}

	wg.Wait()
}
