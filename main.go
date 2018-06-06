package main

import (
	"docviewer/app"
	"docviewer/app/models"
	"docviewer/bootstrap"
	"docviewer/db/gorm"
	"docviewer/db/redis"
	"docviewer/db/mongo"
)

// main entry
func main() {
	// init server
	app.Init()

	// init database
	gorm.Init()
	autoDropTables()
	autoCreateTables()
	autoMigrateTables()

	// init redis
	redis.Init()

	// init mongo
	mongo.Init()

	// run server
	app.Server.Logger.Fatal(app.Server.Start(":1313"))
}

// autoCreateTables: create database tables using GORM
func autoCreateTables() {
	if !gorm.DBManager().HasTable(&models.User{}) {
		gorm.DBManager().CreateTable(&models.User{})
	}
}

// autoMigrateTables: migrate table columns using GORM
func autoMigrateTables() {
	gorm.DBManager().AutoMigrate(&models.User{})
}

// auto drop tables on dev mode
func autoDropTables() {
	if bootstrap.App.ENV == "dev" {
		gorm.DBManager().DropTableIfExists(&models.User{}, &models.User{})
	}
}
