package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Pshimaf-Git/particard/internal/models"
	"github.com/Pshimaf-Git/particard/internal/storage"
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

const DriverName = "sqlite"

type SQLiteStorage struct {
	db *sql.DB
}

// New creates a new SQLite storage instance
func New(ctx context.Context, dataSourceName string) (*SQLiteStorage, error) {
	db, err := sql.Open(DriverName, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("sqlite.New: open database connection: %w", err)
	}

	// Test the connection
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("sqlite.New: ping database: %w", err)
	}

	storage := &SQLiteStorage{db: db}

	// Initialize the database schema
	if err := storage.initSchema(ctx); err != nil {
		return nil, fmt.Errorf("sqlite.New: initialize schema: %w", err)
	}

	return storage, nil
}

// initSchema creates the necessary tables if they don't exist
func (s *SQLiteStorage) initSchema(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS parti_members (
		id BLOB PRIMARY KEY,
		parti TEXT NOT NULL,
		name TEXT NOT NULL,
		role TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_parti_members_parti ON parti_members(parti);
	CREATE INDEX IF NOT EXISTS idx_parti_members_name ON parti_members(name);
	`

	_, err := s.db.ExecContext(ctx, query)
	return err
}

// CreatePartiMember creates a new party member and returns the generated UUID
func (s *SQLiteStorage) CreatePartiMember(ctx context.Context, member *models.PartiMember) (uuid.UUID, error) {
	// Generate UUID if not already set
	if member.ID == uuid.Nil {
		member.ID = uuid.New()
	}

	query := `
	INSERT INTO parti_members (id, parti, name, role)
	VALUES (?, ?, ?, ?)
	`

	_, err := s.db.ExecContext(ctx, query,
		member.ID,
		member.Parti,
		member.Name,
		member.Role,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("sqlite.CreatePartiMember: %w", err)
	}

	return member.ID, nil
}

// GetPartiMember retrieves a party member by UUID
func (s *SQLiteStorage) GetPartiMember(ctx context.Context, uid uuid.UUID) (*models.PartiMember, error) {
	query := `
	SELECT id, parti, name, role
	FROM parti_members 
	WHERE id = ?
	`

	row := s.db.QueryRowContext(ctx, query, uid)

	var member models.PartiMember
	err := row.Scan(
		&member.ID,
		&member.Parti,
		&member.Name,
		&member.Role,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("sqlite.GetPartiMember: %w", storage.ErrPartiMemberNotFound)
		}
		return nil, fmt.Errorf("sqlite.GetPartiMember: %w", err)
	}

	return &member, nil
}

// UpdatePartiMember updates an existing party member
func (s *SQLiteStorage) UpdatePartiMember(ctx context.Context, uid uuid.UUID, newMember *models.PartiMember) error {
	query := `
	UPDATE parti_members 
	SET parti = ?, name = ?, role = ?, updated_at = CURRENT_TIMESTAMP
	WHERE id = ?
	`

	result, err := s.db.ExecContext(ctx, query,
		newMember.Parti,
		newMember.Name,
		newMember.Role,
		uid,
	)
	if err != nil {
		return fmt.Errorf("sqlite.UpdatePartiMember: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("sqlite.UpdatePartiMember: get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("sqlite.UpdatePartiMember: %w", storage.ErrPartiMemberNotFound)
	}

	return nil
}

// RemovePartiMember deletes a party member by UUID
func (s *SQLiteStorage) RemovePartiMember(ctx context.Context, uid uuid.UUID) error {
	query := `DELETE FROM parti_members WHERE id = ?`

	result, err := s.db.ExecContext(ctx, query, uid)
	if err != nil {
		return fmt.Errorf("sqlite.RemovePartiMember: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("sqlite.RemovePartiMember: get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("sqlite.RemovePartiMember: %w", storage.ErrPartiMemberNotFound)
	}

	return nil
}

// Close closes the database connection
func (s *SQLiteStorage) Close() error {
	return s.db.Close()
}
