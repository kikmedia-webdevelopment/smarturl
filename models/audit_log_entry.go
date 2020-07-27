package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/juliankoehn/mchurl/storage"
	"github.com/pkg/errors"
)

// AuditAction holds different action types as a string
type AuditAction string
type auditLogType string

// AuditLogEntry is the database model for audit log entries.
type AuditLogEntry struct {
	ID         uuid.UUID    `json:"id" db:"id"`
	Timestamp  string       `json:"timestamp" db:"timestamp"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	ActorID    uint         `json:"actor_id" db:"actor_id"`
	ActorEmail string       `json:"actor_email" db:"actor_email"`
	Action     AuditAction  `json:"action" db:"action"`
	LogType    auditLogType `json:"log_type" db:"log_type"`
}

const (
	// LoginAction is a action type
	LoginAction AuditAction = "login"
	// LogoutAction is a action type
	LogoutAction AuditAction = "logout"
	// InviteAcceptedAction is a action type
	InviteAcceptedAction AuditAction = "invite_accepted"
	// UserSignedUpAction is a action type
	UserSignedUpAction AuditAction = "user_signedup"
	// UserDeletedAction is a action type
	UserDeletedAction AuditAction = "user_deleted"
	// UserModifiedAction is a action type
	UserModifiedAction AuditAction = "user_modified"
	// UserRecoveryRequestedAction is a action type
	UserRecoveryRequestedAction AuditAction = "user_recovery_requested"
	// TokenRevokedAction is a action type
	TokenRevokedAction AuditAction = "token_revoked"
	// TokenRefreshedAction is a action type
	TokenRefreshedAction AuditAction = "token_refreshed"

	account auditLogType = "account"
	team    auditLogType = "team"
	token   auditLogType = "token"
	user    auditLogType = "user"
)

var actionLogTypeMap = map[AuditAction]auditLogType{
	LoginAction:                 account,
	LogoutAction:                account,
	UserSignedUpAction:          team,
	UserDeletedAction:           team,
	TokenRevokedAction:          token,
	TokenRefreshedAction:        token,
	UserModifiedAction:          user,
	UserRecoveryRequestedAction: user,
}

// TableName returns the audit tablename
func (AuditLogEntry) TableName() string {
	return "audit_log_entries"
}

// NewAuditLogEntry creates a new audit entry
func NewAuditLogEntry(tx *storage.Connection, actor *User, action AuditAction) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return errors.Wrap(err, "Error generating unique id")
	}
	l := AuditLogEntry{
		ID:         id,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		ActorEmail: actor.Email,
		ActorID:    actor.ID,
		Action:     action,
		LogType:    actionLogTypeMap[action],
	}

	if err := tx.Create(&l).Error; err != nil {
		return errors.Wrap(err, "Database error creating audit log entry")
	}

	return nil
}
