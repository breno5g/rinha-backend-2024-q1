package repositories

import (
	"context"

	"github.com/breno5g/rinha-backend-2024-q1/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransacaoRepository interface {
	CreateTransaction(ctx context.Context, transaction entity.Transacao, saldo int64) error
	GetBalance(ctx context.Context, clientId int) (entity.Cliente, error)
	GetTransactions(ctx context.Context, clientId int) ([]entity.Transacao, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateTransaction(ctx context.Context, transaction entity.Transacao, saldo int64) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		"INSERT INTO public.transacoes(tipo, descricao, valor, cliente_id) VALUES ($1, $2, $3, $4)",
		transaction.Tipo, transaction.Descricao, transaction.Valor, transaction.ClienteID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		"UPDATE public.clientes SET saldo = $1 WHERE id = $2",
		saldo, transaction.ClienteID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *repository) GetBalance(ctx context.Context, clientId int) (entity.Cliente, error) {
	var cliente entity.Cliente
	err := r.db.QueryRow(ctx, "SELECT id, nome, saldo, limite FROM public.clientes WHERE id = $1", clientId).Scan(&cliente.ID, &cliente.Nome, &cliente.Saldo, &cliente.Limite)
	if err != nil {
		return entity.Cliente{}, err
	}
	return cliente, nil
}

func (r *repository) GetTransactions(ctx context.Context, clientId int) ([]entity.Transacao, error) {
	rows, err := r.db.Query(ctx, "SELECT id, tipo, descricao, valor, cliente_id, realizada_em FROM public.transacoes WHERE cliente_id = $1 ORDER BY realizada_em DESC LIMIT 5", clientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.Transacao
	for rows.Next() {
		var transaction entity.Transacao
		err = rows.Scan(&transaction.ID, &transaction.Tipo, &transaction.Descricao, &transaction.Valor, &transaction.ClienteID, &transaction.RealizadaEm)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
