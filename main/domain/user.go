package domain

import "time"

type User struct {
	ID                    uint64    `json:"id"`
	Username              string    `json:"username"`
	Password              string    `json:"-"`
	Email                 string    `json:"email"`
	Enabled               bool      `json:"enabled"`
	CreatedAt             time.Time `json:"-"`
	UpdatedAt             time.Time `json:"-"`
	AccountNonExpired     bool      `json:"-"`
	CredentialsNonExpired bool      `json:"-"`
	AccountNonLocked      bool      `json:"-"`
	// TODO Authorities
}
