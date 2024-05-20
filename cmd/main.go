package main

import (
	_ "RolePlayModule/docs"
	"RolePlayModule/internal/di"
	"RolePlayModule/internal/pkg/controllers"
	"RolePlayModule/internal/utils/config"
	"fmt"
)

//	@title			52FOOD API
//	@version		1.0
//	@description	API server for 52FOOD
//@host
//	@BasePath		/

func main() {
	cfg := config.NewConfig()
	cfg.InitENV()

	container := di.New(cfg)
	db := container.GetDB()

	postgresDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get database connection: %v", err))
	}
	if err := postgresDB.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	storage := container.GetSQLStorage()
	server := controllers.NewServer(storage, cfg)
	server.InitSwagger()
	err = server.Run(cfg.ServerPort)
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
