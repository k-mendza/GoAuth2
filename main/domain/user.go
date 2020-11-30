package domain

import "time"

type User struct {
	ID        uint64 `json:"id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Enabled   bool   `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AccountNonExpired     bool `json:"account_non_expired"`
	CredentialsNonExpired bool `json:"credentials_non_expired"`
	AccountNonLocked      bool `json:"account_non_locked"`
	// TODO Authorities
}