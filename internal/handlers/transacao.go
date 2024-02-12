package handlers

import (
	"fmt"
	"strconv"

	"github.com/breno5g/rinha-backend-2024-q1/internal/dto"
	"github.com/breno5g/rinha-backend-2024-q1/internal/entity"
	"github.com/breno5g/rinha-backend-2024-q1/internal/service"
	"github.com/gin-gonic/gin"
)

type TransacaoController struct {
	service service.TransacaoService
}

func NewTransacaoController(service service.TransacaoService) *TransacaoController {
	return &TransacaoController{
		service: service,
	}
}

func (c *TransacaoController) CreateTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := &dto.TransacaoRequest{}
		if err := ctx.ShouldBindJSON(input); err != nil {
			ctx.JSON(422, gin.H{"error": err.Error()})
			return
		}

		if err := input.Validate(); err != nil {
			ctx.JSON(422, gin.H{"error": err.Error()})
			return
		}

		id, _ := strconv.Atoi(ctx.Param("id"))
		cliente, err := c.service.GetBalance(ctx, id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		var newBalance int64

		if input.Tipo == "c" {
			newBalance = cliente.Saldo + input.Valor
		} else {
			newBalance = cliente.Saldo - input.Valor
		}

		if (cliente.Limite + newBalance) < 0 {
			ctx.JSON(422, gin.H{"error": "insufficient funds"})
			return
		}

		fmt.Println(newBalance)
		if newBalance < (cliente.Limite * -1) {
			ctx.JSON(422, gin.H{"error": "insufficient funds"})
			return
		}

		newTransaction := entity.Transacao{
			Tipo:      input.Tipo,
			Descricao: input.Descricao,
			Valor:     input.Valor,
			ClienteID: id,
		}

		err = c.service.CreateTransaction(ctx, newTransaction, newBalance)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		newTransacaoResponse := dto.TransacaoResponse{
			Limite: cliente.Limite,
			Saldo:  newBalance,
		}

		ctx.JSON(200, newTransacaoResponse)
	}
}

func (c *TransacaoController) GetExtract() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		response, err := c.service.GetExtract(ctx, id)
		if err != nil {
			if err.Error() == "no rows in result set" {
				ctx.JSON(404, gin.H{"error": "client not found"})
				return
			}
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, response)
	}
}
