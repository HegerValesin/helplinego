package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserType é um enum para os tipos de usuário
type UserType string

const (
	ADMIN UserType = "ADMIN"
	USER  UserType = "USER"
	SUPER UserType = "SUPER"
	// Adicione outros tipos conforme necessário
)

// User implementa as funcionalidades equivalentes ao modelo Java
type User struct {
	ID         uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name       string     `gorm:"not null" json:"name" binding:"required" validate:"required" form:"name"`
	Login      string     `gorm:"not null;unique" json:"login" binding:"required" validate:"required" form:"login"`
	Email      string     `json:"email" form:"email"`
	Type       UserType   `gorm:"type:varchar(50)" json:"type" form:"type"`
	Password   string     `gorm:"not null" json:"-" binding:"required" validate:"required" form:"password"`
	Facilities []Facility `gorm:"many2many:user_facilities;" json:"facilities"`
	Sector     *Sector    `json:"sector"`
	SectorID   *uuid.UUID `gorm:"type:uuid" json:"sector_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"-"`
}

// BeforeCreate será chamado pelo GORM antes de criar um registro
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

// TableName especifica o nome da tabela
func (User) TableName() string {
	return "users"
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// GetAuthorities retorna as permissões do usuário
func (u *User) GetAuthorities() []string {
	role := "USER"
	if u.Type != "" {
		role = string(u.Type)
	}
	return []string{"ROLE_" + role}
}

// GetUsername retorna o login do usuário
func (u *User) GetUsername() string {
	return u.Login
}

// GetPassword retorna a senha do usuário
func (u *User) GetPassword() string {
	return u.Password
}

// IsAccountNonExpired verifica se a conta não está expirada
func (u *User) IsAccountNonExpired() bool {
	return true
}

// IsAccountNonLocked verifica se a conta não está bloqueada
func (u *User) IsAccountNonLocked() bool {
	return true
}

// IsCredentialsNonExpired verifica se as credenciais não estão expiradas
func (u *User) IsCredentialsNonExpired() bool {
	return true
}

// IsEnabled verifica se a conta está ativa
func (u *User) IsEnabled() bool {
	return true
}
