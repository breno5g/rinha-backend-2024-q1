package dto

import "errors"

type TransacaoRequest struct {
	Valor     int64  `json:"valor" validate:"required,gt=0"`
	Tipo      string `json:"tipo" validate:"required,len=1"`
	Descricao string `json:"descricao" validate:"required,len=10"`
}

type TransacaoResponse struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}

func (t *TransacaoRequest) Validate() error {
	typeValidation := t.Tipo != "c" && t.Tipo != "d"
	descriptionValidation := len(t.Descricao) < 1 || len(t.Descricao) > 10
	valueValidation := t.Valor <= 0
	if typeValidation || descriptionValidation || valueValidation {
		return errors.New("invalid request")
	}

	return nil
}
