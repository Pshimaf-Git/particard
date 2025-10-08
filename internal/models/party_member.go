package models

import "github.com/google/uuid"

type PartiMember struct {
	ID    uuid.UUID `db:"id"`
	Parti string    `db:"parti"`
	Name  string    `db:"name"`
	Role  string    `db:"role"`
}

func NewPartiMember(id uuid.UUID, parti string, name string, role string) *PartiMember {
	return &PartiMember{
		ID:    id,
		Parti: parti,
		Name:  name,
		Role:  role,
	}
}
