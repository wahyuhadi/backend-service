package users

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint           `json:"id"`
	Uuid      uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Username  string         `json:"username"`
	Phone     string         `json:"phone"`
	Email     *string        `json:"email"`
	Password  string         `json:"password"`
	Salt      string         `json:"salt"`
	RoleID    uint           `json:"role_id"`
	CreatedAt time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
	Role      *Role          `json:"role"`
}

type Role struct {
	ID        uint           `json:"id"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
type UserToken struct {
	ID      int64  `json:"id"`
	Role    string `json:"role"`
	Token   string `json:"token"`
	Expired int64  `json:"expired"`
}
