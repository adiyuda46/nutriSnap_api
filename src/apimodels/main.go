package main

import (
	"api_model_cnn/src/apimodels/connection"
	"api_model_cnn/src/apimodels/controller"
	"api_model_cnn/src/apimodels/manager"
	"api_model_cnn/src/apimodels/utils"
	"api_model_cnn/src/config"
	"fmt"
	"net/http"
	"time"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// func init() {
// 	viper.SetConfigFile("src/config/config-dev.json")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// }

func init() {
	viper.SetConfigFile(".env.json")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(fmt.Errorf("error reading config file: %w", err))
	} else {
		fmt.Println("Successfully loaded .env.json configuration.")
	}
}

func main() {
	fmt.Println(viper.GetString("nutrisnap.server"))
	fmt.Println(viper.GetString("token"))

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	addr := ":8000" // port local host nn
	host := "http://192.168.1.8"

	// Router
	router := gin.New()
	// Handle wrong method
	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		utils.HandleError(c, http.StatusMethodNotAllowed, 405, "Method not allowed", "Method Not Allowed", "Method Not Allowed")
	})
	// Handle no route
	router.NoRoute(func(c *gin.Context) {
		utils.HandleError(c, http.StatusNotFound, 404, "Endpoint salah", "Page Not Found", "Page Not Found")
	})
	// init config
	DbConfig := config.CreateConnectionDB()

	// ini connection
	dbConn := connection.CreateConnectionPostgres(DbConfig)

	// init repo
	repo := manager.CreateRepoManager(dbConn)

	// init service
	service := manager.CreateServiceManager(repo)

	// init controller
	controller.CreateNutriSnapController(router, service.NutriSnapService())

	c := cors.AllowAll()
	handler := c.Handler(router)
	newHandler := http.TimeoutHandler(handler, 75*time.Second, "Timeout!")
	//subhost := host[7:]

	router.Use(healthcheck.Default())
	server := &http.Server{
		Addr:         addr,
		Handler:      newHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 75 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	// Start server in goroutine
	go func() {
		// Logging
		logger.Println("Server starting on", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed to start:", err)
		}
	}()

	// Graceful shutdown (optional: add signal handling later)
	logger.Println("Server is running at", host+addr)
	select {} // Keep main alive
}
