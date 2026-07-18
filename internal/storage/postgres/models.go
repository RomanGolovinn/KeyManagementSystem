package postgres

import "time"

type PermissionType string

const (
	None            PermissionType = "none"
	PaymentInitiate PermissionType = "payment:initiate"
	PaymentApprove  PermissionType = "payment:approve"
	PaymentExecute  PermissionType = "payment:execute"
	PaymentReverse  PermissionType = "payment:reverse"

	AccountReadBalance PermissionType = "account:read:balance"
	AccountReadHistory PermissionType = "account:read:history"
	AccountReadPii     PermissionType = "account:read:pii"

	ConfigLimitsWrite PermissionType = "config:limits:write"

	ProductManage  PermissionType = "product:manage"
	SystemOverride PermissionType = "system:override"
)

type Client struct {
	Id      string
	Name    string
	TlsPem  string
	ValidTo time.Time
}

type Server struct {
	Id      string
	Name    string
	Address string
}

type Permission struct {
	Id         int
	ClientId   string
	ServerId   string
	Permission PermissionType
}

type IssuedCredential struct {
	Id               string
	ClientId         string
	ServerId         string
	TempUsername     string
	TempPasswordHash string
	CreatedAt        time.Time
	ValidTo          time.Time
	RevokedAt        time.Time
}

type AuditLog struct {
	Id        string
	Timestamp time.Time
	ClientId  string
	ServerId  string
	EventType string
	IpAddress string
	Details   string // JSONB stored as string
}
