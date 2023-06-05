package entity

import (
	"errors"
	"github.com/yvesromeiro/goapifc/pkg/entity"
)

var (
	errIDIsRequired    = errors.New("id is required")
	errInvalidID       = errors.New("invalid id")
	errNameIsRequired  = errors.New("name is required")
	errPriceIsRequired = errors.New("price is required")
)

type Product struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	CreateAt string    `json:"create_at"`
}

func (p *Product) Validate() error {

}
