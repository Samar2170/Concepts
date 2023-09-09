package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OnboardingConnectionPool() {
	log.Println("OnboardingConnectionPool Initiated")
	dbName := DB_NAME
	dbUser := DB_USER
	dbPass := DB_PASS
	dbTcp := "@tcp(" + DB_HOST + ":" + DB_PORT + ")"
	db, err := gorm.Open(postgres.Open(dbUser+":"+dbPass+dbTcp+"/"+dbName+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Println("Error in OnboardingConnectionPool")
		log.Println(err)
	}
	dbMaxIdleConns := DB_MAX_IDLE_CONNS
	dbMaxOpenConns := DB_MAX_OPEN_CONNS
	dbMaxLifeTime := DB_MAX_LIFE_TIME
	dbMaxIdleTime := DB_MAX_IDLE_TIME

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error in OnboardingConnectionPool")
		log.Println(err)
	}
	sqlDB.SetMaxIdleConns(dbMaxIdleConns)
	sqlDB.SetMaxOpenConns(dbMaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(dbMaxIdleTime) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(dbMaxLifeTime) * time.Second)

	log.Println("OnboardingConnectionPool Completed")
}
