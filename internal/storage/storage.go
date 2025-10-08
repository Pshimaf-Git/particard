package storage

import (
	"context"
	"errors"
	"io"

	"github.com/Pshimaf-Git/particard/internal/models"
	"github.com/google/uuid"
)

var ErrPartiMemberNotFound = errors.New("parti member not found")

type PartiMembersManager interface {
	CreatePartiMember(ctx context.Context, member *models.PartiMember) (uuid.UUID, error)
	GetPartiMember(ctx context.Context, uid uuid.UUID) (*models.PartiMember, error)
	UpdatePartiMember(ctx context.Context, uid uuid.UUID, newMember *models.PartiMember) error
	RemovePartiMember(ctx context.Context, uid uuid.UUID) error
}

type Storage interface {
	PartiMembersManager

	io.Closer
}
