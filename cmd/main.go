package main

import (
	"log"
	"net/http"
	"time"

	"github.com/destroyxiety/CourseWorkSallary/cmd/config"
	"github.com/destroyxiety/CourseWorkSallary/internal/models"
	"github.com/destroyxiety/CourseWorkSallary/internal/repositories"
	"github.com/destroyxiety/CourseWorkSallary/internal/routers"
	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Validator struct {
	V *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.V.Struct(i)
}

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(config.GetDatabaseURL(cfg)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Customers{},
		&models.Positions{},
		&models.Payments{},
		&models.Taxes{},
		&models.PaymentsTaxes{},
		&models.Employees{},
		&models.Deals{},
		&models.Percentages{},
		&models.Accruals{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repoF := repositories.NewRepoFactory(db)
	svcF := services.NewServicesFactory(repoF)

	e := echo.New()
	e.Validator = &Validator{V: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3003"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	routers.RegisterRoutes(e, svcF)

	addr := ":" + cfg.HTTPPort
	srv := &http.Server{
		Addr:         addr,
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("starting server on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
