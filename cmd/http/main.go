package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Oxyaction/go-crud/cmd/http/handlers"

	"github.com/Oxyaction/go-crud/internals/platform/db"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can not read current directory %v", err)
	}
	viper.AddConfigPath(fmt.Sprintf("%s/config", path))
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(viper.GetString("log_level"))
	if err != nil {
		level = log.InfoLevel
	}

	log.SetLevel(level)
}

func main() {
	initConfig()
	initLogger()
	conn, err := db.Open(context.Background(), &db.DbConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Database: viper.GetString("db.database"),
		Username: viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
	})

	defer conn.Close(context.Background())

	if err != nil {
		panic(fmt.Sprintf("Error connecting to db %v", err))
	}

	router := httprouter.New()
	handlers.AttachHandlers(router, conn)

	port := strconv.Itoa(viper.GetInt("port"))
	s := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.WithFields(log.Fields{
		"port": port,
	}).Info("server started")

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
