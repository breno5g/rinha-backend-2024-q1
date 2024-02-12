package service

import (
	"context"
	"errors"
	"time"

	"github.com/breno5g/rinha-backend-2024-q1/internal/entity"
	"github.com/breno5g/rinha-backend-2024-q1/internal/repositories"
)

var (
	ErrClientNotFound = errors.New("cliente not found")
)

type TransacaoService interface {
	CreateTransaction(ctx context.Context, transaction entity.Transacao, saldo int64) error
	GetBalance(ctx context.Context, clientId int) (entity.Cliente, error)
	GetExtract(ctx context.Context, clientId int) (interface{}, error)
}

type transacaoService struct {
	repo repositories.TransacaoRepository
}

func NewService(repo repositories.TransacaoRepository) *transacaoService {
	return &transacaoService{
		repo: repo,
	}
}

func (s *transacaoService) CreateTransaction(ctx context.Context, transaction entity.Transacao, saldo int64) error {
	err := s.repo.CreateTransaction(ctx, transaction, saldo)
	if err != nil {
		return err
	}

	return nil
}

func (s *transacaoService) GetBalance(ctx context.Context, clientId int) (entity.Cliente, error) {
	cliente, err := s.repo.GetBalance(ctx, clientId)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return entity.Cliente{}, ErrClientNotFound
		}
		return entity.Cliente{}, err
	}

	return cliente, nil
}

type Saldo struct {
	Total        int64     `json:"total"`
	Data_extrato time.Time `json:"data_extrato"`
	Limite       int64     `json:"limite"`
}

type ExtractResponse struct {
	Saldo             `json:"saldo"`
	UltimasTransacoes []entity.Transacao `json:"ultimas_transacoes"`
}

func (s *transacaoService) GetExtract(ctx context.Context, clientId int) (interface{}, error) {
	balance, err := s.GetBalance(ctx, clientId)
	if err != nil {
		return nil, err
	}

	transactions, err := s.repo.GetTransactions(ctx, clientId)
	if err != nil {
		return nil, err
	}

	response := ExtractResponse{
		Saldo: Saldo{
			Total:        balance.Saldo,
			Data_extrato: time.Now(),
			Limite:       balance.Limite,
		},
		UltimasTransacoes: transactions,
	}

	return response, nil
}
