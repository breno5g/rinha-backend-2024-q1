package entity

import "time"

type Transacao struct {
	ID          int       `json:"-"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	Valor       int64     `json:"valor"`
	ClienteID   int       `json:"-"`
	RealizadaEm time.Time `json:"realizada_em"`
}
