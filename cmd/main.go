package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	postgres2 "github.com/tranvu1111/go-students-new/internal/infrastructure/db/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/tranvu1111/go-students-new/internal/interface/api/rest"
	"github.com/tranvu1111/go-students-new/internal/application/services"

)

func main(){
	gin.SetMode(gin.ReleaseMode)

	dsn := "host=localhost user=postgres password=tranvu123@ dbname=demodb port=5432 sslmode=disable"
	port := ":8080"
	
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database : %v" , err)
	}

	gormDB.AutoMigrate()

	studentRepo := postgres2.NewGormStudentRepo(gormDB)
	idempotencyRepo := postgres2.NewGormIdempotencyRepository(gormDB)


	studentService := services.NewStudentService(studentRepo, idempotencyRepo)
	

	r := gin.Default()
	rest.NewStudentController(r, studentService)

	
	if err := r.Run(fmt.Sprintf("%s", port));err != nil {
		log.Fatalf("Gin server failed to start: %v", err)
	}
	
}