package audit

import "time"

// Audit represents the audit information for a data entity
type Audit struct {
	CreatedAt        time.Time
	CreatedBy        string
	CreatedOn        string
	CreatedAtDisplay string
	Updates          []AuditUpdateInfo
	DeletedAt        time.Time
	DeletedBy        string
	DeletedOn        string
	DeletedAtDisplay string
	AuditSequence    int
	DBVersion        int
	//Empty     time.Time // Convience Field - Used to avoid erros with dates.
}

type AuditUpdateInfo struct {
	UpdatedAt        time.Time
	UpdateAction     string
	UpdatedBy        string
	UpdatedOn        string
	UpdatedAtDisplay string
	UpdateNotes      string
}

// Action represents an audit action with its properties
type Action struct {
	code        string
	short       string
	description string
	silent      bool
}
