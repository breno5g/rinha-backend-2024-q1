package router

import (
	"github.com/breno5g/rinha-backend-2024-q1/config"
	"github.com/breno5g/rinha-backend-2024-q1/internal/handlers"
	"github.com/breno5g/rinha-backend-2024-q1/internal/repositories"
	"github.com/breno5g/rinha-backend-2024-q1/internal/service"
	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	db := config.GetDB()
	repo := repositories.NewRepository(db)
	service := service.NewService(repo)
	controller := handlers.NewTransacaoController(service)

	r.POST("/clientes/:id/transacoes", controller.CreateTransaction())
	r.GET("/clientes/:id/extrato", controller.GetExtract())

	// apiPort := 9999
	// r.Run(fmt.Sprintf(":%d", apiPort))
	r.Run()
}
